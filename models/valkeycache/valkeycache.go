// Package valkeycache implements the gorm-caches Cacher[1] interface to store
// data queried from Postgres in Valkey as a cache.
//
// [1]: https://pkg.go.dev/github.com/go-gorm/caches/v4#Cacher
package valkeycache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-gorm/caches/v4"
	valkey "github.com/redis/go-redis/v9"
)

// New constructs a new instance of the Valkey cache based on an already
// configured Client.
func New(rdb *valkey.Client) caches.Cacher {
	return &cache{rdb: rdb}
}

type cache struct {
	rdb *valkey.Client
}

func (c *cache) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	res, err := c.rdb.Get(ctx, key).Result()
	if err == valkey.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if err := q.Unmarshal([]byte(res)); err != nil {
		return nil, err
	}

	return q, nil
}

func (c *cache) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.rdb.Set(ctx, key, res, 300*time.Second) // Set proper cache time
	return nil
}

func (c *cache) Invalidate(ctx context.Context) error {
	var (
		cursor uint64
		keys   []string
	)
	for {
		var (
			k   []string
			err error
		)
		k, cursor, err = c.rdb.Scan(ctx, cursor, fmt.Sprintf("%s*", caches.IdentifierPrefix), 0).Result()
		if err != nil {
			return err
		}
		keys = append(keys, k...)
		if cursor == 0 {
			break
		}
	}

	if len(keys) > 0 {
		if _, err := c.rdb.Del(ctx, keys...).Result(); err != nil {
			return err
		}
	}
	return nil
}
