package main

import (
	"database/sql"
	"log"

	"github.com/bdtomlin/gator/internal/config"
	"github.com/bdtomlin/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func newState() (*state, error) {
	cfg, err := config.Read()
	if err != nil {
		return &state{}, err
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	return &state{
		cfg: cfg,
		db:  dbQueries,
	}, nil
}
