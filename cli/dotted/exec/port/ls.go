/*
  exec port.ls

  @author Mikiya Michishita
*/

package port

import (
	"context"
	"flag"
	"github.com/go-template/cli/dotted/exec/cli"
	"github.com/go-template/cli/dotted/exec/flags"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	feature string
}

func init() {
	cli.Register("port.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag = flags.NewClientFlag()
}

func (cmd *ls) Process(ctx context.Context) error {
	var err error
	return err
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	var err error
	return err
}
