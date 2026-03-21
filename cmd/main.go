package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	dbconfig := DBConfig{}
	cfg := Config{
		dbConfig: dbconfig,
		addr:     ":8089",
	}

	api := Application{
		config: cfg,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	if err := api.run(api.mount()); err != nil {
		log.Fatal(err.Error())

	}
}
