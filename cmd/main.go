package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	dbconfig := DBConfig{
		dsn: "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable",
	}
	cfg := Config{
		dbConfig: dbconfig,
		addr:     ":8089",
	}

	conn, err := pgx.Connect(ctx, cfg.dbConfig.dsn)
	if err != nil {
		panic(err)
	}
	api := Application{
		config: cfg,
		db:     conn,
	}
	defer conn.Close(ctx)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	if err := api.run(api.mount()); err != nil {
		log.Fatal(err.Error())

	}
}
