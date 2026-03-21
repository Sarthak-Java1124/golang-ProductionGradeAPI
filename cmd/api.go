package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	config Config
}

type Config struct {
	addr     string
	dbConfig DBConfig
}

func (app *Application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(time.Second * 60))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All Good !!"))
	})

	return r
}

func (app *Application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("The server has started at %s ", app.config.addr)
	return srv.ListenAndServe()
}

type DBConfig struct {
	dsn string
}
