package main

import (
	"log"
	"os"

	"github.com/aizu-go-kapro/dir"
)

const (
	COMMAND = "echo"
	ARG     = "HELLO"
)

func main() {
	executor := dir.NewExecutor(
		dir.SetCommand(COMMAND),
		dir.SetArgs(ARG),
	)
	if err := executor.Do(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
