/*
  Go Project Description

  @author Mikiya Michishita
*/

package cliobjs

import (
	"github.com/go-template/cli/dotted/cliobjs/types"
)

type Client struct {
	Content types.Content
}

func NewClient() (*Client) {
	c := Client {}
	// c.Content = 
	return &c, nil
}
