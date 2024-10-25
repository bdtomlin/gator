package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/bdtomlin/gator/internal/config"
	"github.com/bdtomlin/gator/internal/database"
	_ "github.com/lib/pq"
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

func main() {
	st, err := newState()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmds := newCommands()

	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)

	cmd, err := newCommand(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = cmds.run(st, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
