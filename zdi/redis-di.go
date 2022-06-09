package zdi

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jigadhirasu/zzz/z"
)

type RedisDI func(dbr *redis.Client) z.Bytes

type RedisCtxDI func(dbr *redis.Client, ctx context.Context) z.Bytes
