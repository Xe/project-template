package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Xe/project-template/models/modelstest"
)

func makeServer(t *testing.T) (*Server, *httptest.Server) {
	t.Helper()

	dbURL := modelstest.MaybeSpawnDB(t)
	redisURL := modelstest.MaybeSpawnValkey(t)

	s, err := New(Options{
		DatabaseURL: dbURL,
		RedisURL:    redisURL,
	})
	if err != nil {
		t.Fatal(err)
	}

	mux := http.NewServeMux()
	s.register(mux)

	ts := httptest.NewServer(mux)
	t.Cleanup(func() {
		ts.Close()
	})

	return s, ts
}

func TestNew(t *testing.T) {
	makeServer(t)
}
