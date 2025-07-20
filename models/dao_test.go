package models

import (
	"testing"

	"github.com/Xe/project-template/models/modelstest"
)

func TestNewDAO(t *testing.T) {
	dbURL := modelstest.MaybeSpawnDB(t)
	redisURL := modelstest.MaybeSpawnValkey(t)

	rdb, err := ConnectValkey(redisURL)
	if err != nil {
		t.Fatal(err)
	}

	dao, err := New(dbURL, rdb)
	if err != nil {
		t.Fatal(err)
	}
	_ = dao
}
