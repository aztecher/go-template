/*
  Go Project Description

  Date :
  @author Mikiya Michishita
*/

package flags

import "sync"

type common struct {
	register sync.Once
	process  sync.Once
}

func (c *common) RegisterOnce(fn func()) {
	c.register.Do(fn)
}

func (c *common) ProcessOnce(fn func() error) (err error) {
	c.process.Do(func() {
		err = fn()
	})
	return err
}
