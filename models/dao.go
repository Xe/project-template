package models

import (
	"fmt"
	"log/slog"
	"time"

	slogGorm "github.com/orandin/slog-gorm"
	valkey "github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DAO struct {
	db *gorm.DB
}

func (d *DAO) DB() *gorm.DB {
	return d.db
}

func New(dbURL string, rdb *valkey.Client) (*DAO, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: slogGorm.New(
			slogGorm.WithErrorField("err"),
			slogGorm.WithRecordNotFoundError(),
			slogGorm.SetLogLevel(slogGorm.DefaultLogType, slog.LevelDebug),
		),
	})
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	// Use advisory lock to prevent concurrent migrations with exponential backoff
	var lockResult bool
	backoff := 125 * time.Millisecond
	maxRetries := 5

	for attempt := 0; attempt < maxRetries; attempt++ {
		err = db.Raw("SELECT pg_try_advisory_lock(12345)").Scan(&lockResult).Error
		if err != nil {
			return nil, fmt.Errorf("can't acquire migration lock: %w", err)
		}

		if lockResult {
			break
		}

		// If not the last attempt, wait with exponential backoff
		if attempt < maxRetries-1 {
			time.Sleep(backoff)
			backoff *= 2 // Double the backoff time for next attempt
		}
	}

	if !lockResult {
		return nil, fmt.Errorf("failed to acquire migration lock after %d attempts", maxRetries)
	}

	defer func() {
		db.Raw("SELECT pg_advisory_unlock(12345)")
	}()

	// Now run migration safely
	err = db.AutoMigrate(&Example{})
	if err != nil {
		return nil, fmt.Errorf("can't run migrations: %w", err)
	}

	// cachesPlugin := &caches.Caches{Conf: &caches.Config{
	// 	Easer:  true,
	// 	Cacher: valkeycache.New(rdb),
	// }}

	// if err := db.Use(cachesPlugin); err != nil {
	// 	return nil, fmt.Errorf("can't configure cache: %w", err)
	// }

	return &DAO{
		db: db,
	}, nil
}
