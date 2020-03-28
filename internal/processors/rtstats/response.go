package rtstats

import (
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	d2pb "github.com/paralin/go-dota2/protocol"
	"golang.org/x/xerrors"
	"google.golang.org/protobuf/encoding/protojson"

	nsjson "github.com/13k/night-stalker/internal/json"
)

type response struct {
	*resty.Response
}

func (r *response) Parse() (*d2pb.CMsgDOTARealtimeGameStatsTerse, error) {
	b, err := r.transform()

	if err != nil {
		return nil, xerrors.Errorf("error transforming response: %w", err)
	}

	if b == nil {
		return nil, nil
	}

	pbmsg := &d2pb.CMsgDOTARealtimeGameStatsTerse{}

	if err := protojson.Unmarshal(b, pbmsg); err != nil {
		return nil, xerrors.Errorf("error unmarshaling transformed response: %w", err)
	}

	return pbmsg, nil
}

func (r *response) transform() ([]byte, error) {
	t, err := nsjson.NewTransform(r.Body())

	if err != nil {
		return nil, xerrors.Errorf("error creating JSON transform: %w", err)
	}

	if t.ReadNil() {
		return nil, nil
	}

	// root
	b, err := t.Object(func(key string) ([]byte, error) {
		if key != "teams" {
			return t.SkipAndReturnBytes(), nil
		}

		// ignore non-array "teams" field
		if typ := t.WhatIsNext(); typ != jsoniter.ArrayValue {
			return nil, nil
		}

		// teams array
		return t.Array(func() ([]byte, error) {
			// ignore non-object teams items
			if typ := t.WhatIsNext(); typ != jsoniter.ObjectValue {
				return nil, nil
			}

			// team object
			return t.Object(func(key string) ([]byte, error) {
				if key != "players" {
					return t.SkipAndReturnBytes(), nil
				}

				// ignore non-array "players" field
				if typ := t.WhatIsNext(); typ != jsoniter.ArrayValue {
					return nil, nil
				}

				// players array
				return t.Array(func() ([]byte, error) {
					// skip bogus player (usually an empty array)
					if typ := t.WhatIsNext(); typ != jsoniter.ObjectValue {
						return nil, nil
					}

					return t.SkipAndReturnBytes(), nil
				})
			})
		})
	})

	if err != nil {
		return nil, xerrors.Errorf("error transforming response: %w", err)
	}

	if !jsoniter.Valid(b) {
		return nil, xerrors.New("transform() generated invalid JSON")
	}

	return b, nil
}
