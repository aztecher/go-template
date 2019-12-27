package flags

import (
	"sync"
	"flag"
)

type HasFlags interface {
	Register(f *flag.FlagSet)
}

type FlagCommon struct {
	register sync.Once
}

func (fc *FlagCommon) RegisterOnce(f func()) {
	fc.register.Do(f)
}
