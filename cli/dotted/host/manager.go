/*
  Go Project Description

  Date :
  @author Mikiya Michishita
*/

package host

import (
	"github.com/go-template/cli/dotted/object"
)

type Manager struct {
	object.Common
}

func NewManager() *Manager {
	m := Manager {
		object.NewCommon(),
	}
	return &m
}


