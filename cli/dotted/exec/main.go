/*
  exec command line entry point

  @author Mikiya Michishita
*/

package main

import (
	"os"
	"github.com/go-template/cli/dotted/exec/cli"
	_ "github.com/go-template/cli/dotted/exec/host"
	_ "github.com/go-template/cli/dotted/exec/port"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
