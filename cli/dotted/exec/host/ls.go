/*
  exec host.ls

  @author Mikiya Michishita
*/

package host

import (
	"context"
	"flag"
	"github.com/go-template/cli/dotted/exec/cli"
	"github.com/go-template/cli/dotted/exec/flags"
	// "github.com/go-template/cli/dotted/host"
)

type ls struct {
	*flags.ClientFlag
	*flags.OutputFlag

	feature string
}

func init() {
	cli.Register("host.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
	cmd.ClientFlag = flags.NewClientFlag()
	cmd.ClientFlag.Register(ctx, f)

	cmd.OutputFlag = flags.NewOutputFlag()
	cmd.OutputFlag.Register(ctx, f)
}

func (cmd *ls) Process(ctx context.Context) error {
	if err := cmd.ClientFlag.Process(ctx); err != nil {
		return err
	}
	if err := cmd.OutputFlag.Process(ctx); err != nil {
		return err
	}
	return nil
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	var err error
	// client -> manager -> process

	// h := host.NewManager()
	return err
}
