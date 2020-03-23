package context

import (
	"context"

	"github.com/13k/geyser"
	geyserd2 "github.com/13k/geyser/dota2"
	"github.com/faceit/go-steam"
	"github.com/jinzhu/gorm"
	"github.com/paralin/go-dota2"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nslog "github.com/13k/night-stalker/internal/logger"
	nsrds "github.com/13k/night-stalker/internal/redis"
)

type ctxKey int

const (
	ctxKeyLogger ctxKey = iota
	ctxKeyDB
	ctxKeyRedis
	ctxKeySteam
	ctxKeyDota
	ctxKeyBus
	ctxKeyAPI
	ctxKeyDotaAPI
)

func WithLogger(ctx context.Context, logger *nslog.Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger, logger)
}

func GetLogger(ctx context.Context) *nslog.Logger {
	i := ctx.Value(ctxKeyLogger)
	v, ok := i.(*nslog.Logger)

	if !ok {
		return nil
	}

	return v
}

func WithDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxKeyDB, db)
}

func GetDB(ctx context.Context) *gorm.DB {
	i := ctx.Value(ctxKeyDB)
	v, ok := i.(*gorm.DB)

	if !ok {
		return nil
	}

	return v
}

func WithRedis(ctx context.Context, rds *nsrds.Redis) context.Context {
	return context.WithValue(ctx, ctxKeyRedis, rds)
}

func GetRedis(ctx context.Context) *nsrds.Redis {
	i := ctx.Value(ctxKeyRedis)
	v, ok := i.(*nsrds.Redis)

	if !ok {
		return nil
	}

	return v
}

func WithSteam(ctx context.Context, steam *steam.Client) context.Context {
	return context.WithValue(ctx, ctxKeySteam, steam)
}

func GetSteam(ctx context.Context) *steam.Client {
	i := ctx.Value(ctxKeySteam)
	v, ok := i.(*steam.Client)

	if !ok {
		return nil
	}

	return v
}

func WithDota(ctx context.Context, dota *dota2.Dota2) context.Context {
	return context.WithValue(ctx, ctxKeyDota, dota)
}

func GetDota(ctx context.Context) *dota2.Dota2 {
	i := ctx.Value(ctxKeyDota)
	v, ok := i.(*dota2.Dota2)

	if !ok {
		return nil
	}

	return v
}

func WithBus(ctx context.Context, bus *nsbus.Bus) context.Context {
	return context.WithValue(ctx, ctxKeyBus, bus)
}

func GetBus(ctx context.Context) *nsbus.Bus {
	i := ctx.Value(ctxKeyBus)
	v, ok := i.(*nsbus.Bus)

	if !ok {
		return nil
	}

	return v
}

func WithAPI(ctx context.Context, api *geyser.Client) context.Context {
	return context.WithValue(ctx, ctxKeyAPI, api)
}

func GetAPI(ctx context.Context) *geyser.Client {
	i := ctx.Value(ctxKeyAPI)
	v, ok := i.(*geyser.Client)

	if !ok {
		return nil
	}

	return v
}

func WithDotaAPI(ctx context.Context, api *geyserd2.Client) context.Context {
	return context.WithValue(ctx, ctxKeyDotaAPI, api)
}

func GetDotaAPI(ctx context.Context) *geyserd2.Client {
	i := ctx.Value(ctxKeyDotaAPI)
	v, ok := i.(*geyserd2.Client)

	if !ok {
		return nil
	}

	return v
}
