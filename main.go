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

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
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
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAgg)

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
