package main

import (
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	st, err := newState()
	if err != nil {
		log.Fatal(err)
	}

	cmds := newCommands()

	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", middlewareLoggedIn(handleFollow))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))
	cmds.register("following", middlewareLoggedIn(handleFollowing))

	cmd, err := newCommand(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	err = cmds.run(st, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
