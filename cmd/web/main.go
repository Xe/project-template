package main

import (
	"context"
	"flag"
	"log"
	"log/slog"
	"net/http"

	"git.xeserv.us/xe/project-template/internal"
	"git.xeserv.us/xe/project-template/models"
	"github.com/facebookgo/flagenv"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rbcervilla/redisstore/v9"
)

var (
	bind        = flag.String("bind", ":4000", "TCP host:port to bind on")
	databaseURL = flag.String("database-url", "", "Postgres database URL")
	redisURL    = flag.String("redis-url", "", "Valkey/Redis database URL")
	slogLevel   = flag.String("slog-level", "INFO", "log level")
)

func main() {
	flagenv.Parse()
	flag.Parse()

	internal.InitSlog(*slogLevel)

	rdb, err := models.ConnectValkey(*redisURL)
	if err != nil {
		log.Fatal(err)
	}

	dao, err := models.New(*databaseURL, rdb)
	if err != nil {
		log.Fatal(err)
	}

	store, err := redisstore.NewRedisStore(context.Background(), rdb)
	if err != nil {
		log.Fatal(err)
	}

	store.KeyPrefix("session:")
	store.Options(sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400 * 60,
	})

	s := &Server{
		dao:   dao,
		store: store,
	}

	mux := http.NewServeMux()
	s.register(mux)

	slog.Info("now listening", "url", "http://localhost"+*bind)
	log.Fatal(http.ListenAndServe(*bind, mux))
}
