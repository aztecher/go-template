/*
  Go Project Description

  Date :
  @author Mikiya Michishita
*/

package cli

import (
	"context"
	"flag"
	"os"
	"io"
	"io/ioutil"
	"fmt"
	"strings"
	"sort"
)

type HasFlags interface {
	Register(ctx context.Context, f *flag.FlagSet)

	Process(ctx context.Context) error
}

type Command interface {
	HasFlags

	Run(ctx context.Context, f *flag.FlagSet) error
}

func generalHelp(w io.Writer, filter string) {
	var cmds, matches []string
	for name := range commands {
		cmds = append(cmds, name)

		if filter != "" && strings.Contains(name, filter) {
			matches = append(matches, name)
		}
	}

	if len(matches) == 0 {
		fmt.Fprintf(w, "Usage of %s:\n", os.Args[0])
	} else {
		fmt.Fprintf(w, "%s: command '%s' not found, did you mean:\n", os.Args[0], filter)
		cmds = matches
	}

	sort.Strings(cmds)
	for _, name := range cmds {
		fmt.Fprintf(w, "  %s\n", name)
	}
}

func Run(args []string) int {
	hw := os.Stderr
	rc := 1
	hwrc := func(arg string) {
		if arg == "-h" {
			hw = os.Stdout
			rc = 0
		}
	}

	var err error
	if len(args) == 0 {
		generalHelp(hw, "")
		return rc
	}

	name := args[0]
	cmd, ok := commands[name]
	if !ok {
		hwrc(name)
		generalHelp(hw, name)
		return rc
	}

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)

	ctx := context.Background()

	cmd.Register(ctx, fs)

	// flag parse
	if err = fs.Parse(args[1:]); err != nil {
		goto error
	}

	if err = cmd.Process(ctx); err != nil {
		goto error
	}

	if err = cmd.Run(ctx, fs); err != nil {
		goto error
	}

	return 0

error:
	if err == flag.ErrHelp {
		if len(args) == 2 {
			hwrc(args[1])
		}
		// commandHelp
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[0], err)
	}

	return rc
}
