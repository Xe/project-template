package models

import (
	"fmt"
	"log/slog"

	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DAO struct {
	db *gorm.DB
}

func (d *DAO) DB() *gorm.DB {
	return d.db
}

func New(dbURL string) (*DAO, error) {
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

	if err := db.AutoMigrate(
		// put list of models here
		&Example{},
	); err != nil {
		return nil, fmt.Errorf("can't run migrations: %w", err)
	}

	return &DAO{
		db: db,
	}, nil
}
