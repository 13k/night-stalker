package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"golang.org/x/xerrors"

	nslog "github.com/13k/night-stalker/internal/logger"
	nsm "github.com/13k/night-stalker/models"
)

type ModelPersistence interface {
	Find(ctx context.Context, record nsm.Record, query *SelectQuery) (exists bool, err error)
	FindBy(ctx context.Context, record nsm.Record, col string, colValue interface{}) (exists bool, err error)
	FindByID(ctx context.Context, record nsm.Record, id nsm.ID) (exists bool, err error)
	FindAll(ctx context.Context, model nsm.Model, query *SelectQuery, dst interface{}) error
	Filter(ctx context.Context, model nsm.Model, filter SelectQueryFilter, dst interface{}) error
	PluckCol(ctx context.Context, model nsm.Model, col string, query *SelectQuery, dst interface{}) error
	PluckID(ctx context.Context, model nsm.Model, query *SelectQuery) ([]nsm.ID, error)
	RandomID(ctx context.Context, model nsm.Model) (nsm.ID, error)
	Random(ctx context.Context, record nsm.Record) (exists bool, err error)
	Create(ctx context.Context, record nsm.Record) error
	Update(ctx context.Context, record nsm.Record) error
	Upsert(ctx context.Context, record nsm.Record, query *SelectQuery) (created bool, err error)
	Eagerload(ctx context.Context, assoc string, sources ...nsm.Record) error
	EagerloadAssoc(ctx context.Context, assoc nsm.Association, sources ...nsm.Record) error
}

func NewModelPersistence(qb QueryBuilder, qx QueryExecutor, log *nslog.Logger) ModelPersistence {
	return &modelPersistence{qb: qb, qx: qx, log: log}
}

var _ ModelPersistence = (*modelPersistence)(nil)

type modelPersistence struct {
	qb  QueryBuilder
	qx  QueryExecutor
	log *nslog.Logger
}

func (ma *modelPersistence) begin(ctx context.Context, options *sql.TxOptions) (ModelPersistence, QueryTx, error) {
	qtx, txerr := ma.qx.Begin(ctx, options)

	if txerr != nil {
		return nil, nil, xerrors.Errorf("error opening transaction: %w", txerr)
	}

	mq := NewModelPersistence(ma.qb, qtx, ma.log)

	return mq, qtx, nil
}

func (ma *modelPersistence) Find(ctx context.Context, record nsm.Record, q *SelectQuery) (bool, error) {
	q = q.
		From(record.Model().Table()).
		Limit(1).
		Trace()

	exists, err := ma.qx.ScanStruct(ctx, q, record)

	if err != nil {
		return false, xerrors.Errorf("error querying database: %w", err)
	}

	return exists, nil
}

func (ma *modelPersistence) FindBy(ctx context.Context, record nsm.Record, col string, v interface{}) (bool, error) {
	t := record.Model().Table()
	q := ma.qb.
		Select(t.All()).
		Eq(t.Col(col), v).
		Trace()

	exists, err := ma.Find(ctx, record, q)

	if err != nil {
		return false, xerrors.Errorf("error finding model record: %w", err)
	}

	return exists, nil
}

func (ma *modelPersistence) FindByID(ctx context.Context, record nsm.Record, id nsm.ID) (bool, error) {
	t := record.Model().Table()
	q := ma.qb.
		Select(t.All()).
		Eq(t.PK(), id).
		Trace()

	exists, err := ma.Find(ctx, record, q)

	if err != nil {
		return false, xerrors.Errorf("error finding model record: %w", err)
	}

	return exists, nil
}

func (ma *modelPersistence) FindAll(ctx context.Context, model nsm.Model, q *SelectQuery, dst interface{}) error {
	t := model.Table()
	q = q.Select(t.All()).From(t).Trace()

	if err := ma.qx.ScanStructs(ctx, q, dst); err != nil {
		return xerrors.Errorf("error querying database: %w", err)
	}

	return nil
}

func (ma *modelPersistence) Filter(
	ctx context.Context,
	model nsm.Model,
	filter SelectQueryFilter,
	dst interface{},
) error {
	q := ma.qb.Select().Filter(filter)

	if err := ma.FindAll(ctx, model, q, dst); err != nil {
		return xerrors.Errorf("error finding model records: %w", err)
	}

	return nil
}

func (ma *modelPersistence) PluckID(ctx context.Context, model nsm.Model, q *SelectQuery) ([]nsm.ID, error) {
	var ids []nsm.ID

	q = q.
		Select(model.Table().PK()).
		From(model.Table()).
		Trace()

	if err := ma.qx.ScanVals(ctx, q, &ids); err != nil {
		return nil, xerrors.Errorf("error querying database: %w", err)
	}

	return ids, nil
}

func (ma *modelPersistence) PluckCol(
	ctx context.Context,
	model nsm.Model,
	col string,
	q *SelectQuery,
	dst interface{},
) error {
	t := model.Table()
	q = q.Select(t.Col(col)).From(t).Trace()

	if err := ma.qx.ScanVals(ctx, q, dst); err != nil {
		return xerrors.Errorf("error querying database: %w", err)
	}

	return nil
}

// RandomID returns a random ID in the range (MIN(id), MAX(id)) from the model's table.
//
// It returns 0 if there are not records in the table.
//
// NOTE: It doesn't check if a row with the generated ID actually exists in the model's table, so it
// can generate "invalid" IDs. The client is responsible for handling this case (e.g., re-query
// until a valid ID is found).
func (ma *modelPersistence) RandomID(ctx context.Context, model nsm.Model) (nsm.ID, error) {
	pk := model.Table().PK()

	// GREATEST(MIN(id) + ROUND(RANDOM() * MAX(id)), 0) AS id
	c := goqu.Func(
		"GREATEST",
		goqu.L(
			"? + ?",
			goqu.MIN(pk),
			goqu.Func("ROUND", goqu.L(
				"? * ?",
				goqu.Func("RANDOM"),
				goqu.MAX(pk),
			)),
		),
		0,
	).As("id")

	q := ma.qb.
		Select(c).
		From(model.Table()).
		Trace()

	var id nsm.ID

	_, err := ma.qx.ScanVal(ctx, q, &id)

	if err != nil {
		return 0, xerrors.Errorf("error querying database: %w", err)
	}

	return id, nil
}

// Random queries for a random model record.
//
// If there are no records in the model's table, it will return `(false, nil)`
//
// NOTE: It uses `RandomID` to generate a random ID. If it got a missing ID, it will return
// `(false, nil)` and the client is responsible for handling this case.
func (ma *modelPersistence) Random(ctx context.Context, record nsm.Record) (bool, error) {
	id, err := ma.RandomID(ctx, record.Model())

	if err != nil {
		return false, xerrors.Errorf("error fetching random ID: %w", err)
	}

	if id == 0 {
		return false, nil
	}

	exists, err := ma.FindByID(ctx, record, id)

	if err != nil {
		return false, xerrors.Errorf("error finding model record: %w", err)
	}

	return exists, nil
}

func (ma *modelPersistence) Create(ctx context.Context, record nsm.Record) error {
	_, tx, txerr := ma.begin(ctx, nil)

	if txerr != nil {
		return xerrors.Errorf("error opening transaction: %w", txerr)
	}

	if tr, ok := record.(nsm.Timestampable); ok {
		now := time.Now()

		if tr.GetCreatedAt().IsZero() {
			tr.SetCreatedAt(now)
		}

		if tr.GetUpdatedAt().IsZero() {
			tr.SetUpdatedAt(now)
		}
	}

	if err := nsm.BeforeCreate(record); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return xerrors.Errorf("error running before_create callbacks: %w", err)
	}

	t := record.Model().Table()
	q := ma.qb.
		Insert(t).
		Rows(record).
		Returning(t.PK()).
		Trace()

	if _, err := tx.ScanStruct(ctx, q, record); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return xerrors.Errorf("error querying database: %w", err)
	}

	if err := nsm.AfterCreate(record); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return xerrors.Errorf("error running after_create callbacks: %w", err)
	}

	if txerr := tx.Commit(); txerr != nil {
		return xerrors.Errorf("error committing transaction: %w", txerr)
	}

	return nil
}

func (ma *modelPersistence) Update(ctx context.Context, record nsm.Record) error {
	_, tx, txerr := ma.begin(ctx, nil)

	if txerr != nil {
		return xerrors.Errorf("error opening transaction: %w", txerr)
	}

	if tr, ok := record.(nsm.Timestampable); ok {
		tr.SetUpdatedAt(time.Now())
	}

	if err := nsm.BeforeUpdate(record); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return xerrors.Errorf("error running before_update callbacks: %w", err)
	}

	t := record.Model().Table()
	q := ma.qb.
		Update(t).
		Set(record).
		Where(t.PK().Eq(record.GetID())).
		Trace()

	_, err := tx.Exec(ctx, q)

	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return xerrors.Errorf("error executing query: %w", err)
	}

	if err := nsm.AfterUpdate(record); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return xerrors.Errorf("error running after_update callbacks: %w", err)
	}

	if txerr := tx.Commit(); txerr != nil {
		return xerrors.Errorf("error committing transaction: %w", txerr)
	}

	return nil
}

func (ma *modelPersistence) Upsert(ctx context.Context, record nsm.Record, q *SelectQuery) (bool, error) {
	mtx, tx, txerr := ma.begin(ctx, nil)

	if txerr != nil {
		return false, xerrors.Errorf("error opening transaction: %w", txerr)
	}

	rr := record.Model().NewRecord()
	q = q.Select(record.Model().Table().All())

	exists, err := mtx.Find(ctx, rr, q)

	if err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return false, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return false, xerrors.Errorf("error finding model: %w", err)
	}

	if exists {
		changed, err := rr.AssignPartialRecord(record)

		if err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return false, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return false, xerrors.Errorf("error assigning changes: %w", err)
		}

		if changed {
			if err := mtx.Update(ctx, rr); err != nil {
				if txerr := tx.Rollback(); txerr != nil {
					return false, xerrors.Errorf("error rolling back transaction: %w", txerr)
				}

				return false, xerrors.Errorf("error updating model record: %w", err)
			}
		}

		if _, err := record.AssignRecord(rr); err != nil {
			if txerr := tx.Rollback(); txerr != nil {
				return false, xerrors.Errorf("error rolling back transaction: %w", txerr)
			}

			return false, xerrors.Errorf("error assigning changes: %w", err)
		}

		if txerr := tx.Commit(); txerr != nil {
			return false, xerrors.Errorf("error committing transaction: %w", txerr)
		}

		return false, nil
	}

	if err := mtx.Create(ctx, record); err != nil {
		if txerr := tx.Rollback(); txerr != nil {
			return false, xerrors.Errorf("error rolling back transaction: %w", txerr)
		}

		return false, xerrors.Errorf("error creating model record: %w", err)
	}

	if txerr := tx.Commit(); txerr != nil {
		return false, xerrors.Errorf("error committing transaction: %w", txerr)
	}

	return true, nil
}

/*
# eagerload

1. ids = collect IDs
 * belongs_to: source.FK [corresponding to assoc.FK()]
 * has_one:    source.PK [corresponding to assoc.PK()]
 * has_many:   source.PK [corresponding to assoc.PK()]
2. col = determine association column
 * belongs_to: dest.PK [assoc.PK()]
 * has_one:    dest.FK [assoc.FK()]
 * has_many:   dest.FK [assoc.FK()]
3. q = match association column with collected IDs
 * all: In(col, ids)
4. associated = load associated records
 * all: destModel.NewSlicePtr; FindModels(destModel, q, slicePtr); destModel.AsRecordSlice
5. groupBy = transform associated records
 * belongs_to: group by PK
 * has_one:    group by FK
 * has_many:   group by FK
6. mapped = map source to associated
 * belongs_to: mapped[source] = groupBy[source.FK]
 * has_one:    mapped[source] = groupBy[source.PK]
 * has_many:   mapped[source] = groupBy[source.PK]
7. set associated records
 * all: source.SetAssociated(srcName, mapped[source]...)
*/

func (ma *modelPersistence) EagerloadAssoc(ctx context.Context, assoc nsm.Association, sources ...nsm.Record) error {
	if len(sources) == 0 {
		return nil
	}

	model := sources[0].Model()

	for _, r := range sources[1:] {
		if r.Model().Name() != model.Name() {
			return xerrors.Errorf("invalid sources: %w", &ErrMixedRecords{
				ExpectedModel: model,
				InvalidModel:  r.Model(),
			})
		}
	}

	assocIDs, err := assoc.CollectIDs(sources)

	if err != nil {
		return xerrors.Errorf("error collecting IDs: %w", err)
	}

	mDest := assoc.Dest().Model
	slicePtr := mDest.NewSlicePtr()

	q := ma.qb.
		Select().
		In(assoc.Col(), assocIDs).
		Prepared(true).
		Trace()

	if err = ma.FindAll(ctx, mDest, q, slicePtr); err != nil {
		return xerrors.Errorf("error finding model records: %w", err)
	}

	associated, err := mDest.AsRecordSlice(slicePtr)

	if err != nil {
		return xerrors.Errorf("error converting records slice: %w", err)
	}

	if err := assoc.SetRecords(sources, associated); err != nil {
		return xerrors.Errorf("error associating records: %w", err)
	}

	return nil
}

func (ma *modelPersistence) Eagerload(ctx context.Context, assocName string, sources ...nsm.Record) error {
	if len(sources) == 0 {
		return nil
	}

	assoc, err := sources[0].Model().Association(assocName)

	if err != nil {
		return xerrors.Errorf("error finding model association: %w", err)
	}

	if err := ma.EagerloadAssoc(ctx, assoc, sources...); err != nil {
		return xerrors.Errorf("error eagerloading association: %w", err)
	}

	return nil
}
