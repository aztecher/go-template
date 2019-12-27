package main

import (
	"os"
	"github.com/go-template/cli/action-resource/cli"
	_ "github.com/go-template/cli/action-resource/action"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
