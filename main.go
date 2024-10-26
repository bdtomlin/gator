package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

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
	cmds.register("addfeed", handleAddFeed)

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
