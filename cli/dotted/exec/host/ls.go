/*
  exec host.ls

  @author Mikiya Michishita
*/

package host

import (
	"context"
	"flag"
	"github.com/go-template/cli/dotted/exec/cli"
)

type ls struct {

	feature string
}

func init() {
	cli.Register("host.ls", &ls{})
}

func (cmd *ls) Register(ctx context.Context, f *flag.FlagSet) {
}

func (cmd *ls) Process(ctx context.Context) error {
	var err error
	return err
}

func (cmd *ls) Run(ctx context.Context, f *flag.FlagSet) error {
	var err error
	return err
}
