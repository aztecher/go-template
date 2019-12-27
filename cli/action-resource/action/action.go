package action

import (
	"github.com/go-template/cli/action-resource/flags"
	"github.com/go-template/cli/action-resource/resource"
)

var actions = ActionMap{}

type ActionMap []NamedAction

func (am *ActionMap) Register(name string, a Action) {
	na := NamedAction{ Name: name, Action: a}
	*am = append(*am, na)
}

func (am *ActionMap) Search(args []string) (Action, bool) {
	for _, namedAction := range *am {
		_, ok := namedAction.search(args)
		if ok {
			return namedAction.Action, ok
		}
	}
	return nil, false
}

type NamedAction struct {
	Name   string
	Action Action
}

func (na *NamedAction) search(args []string) (string, bool) {
	for _, arg := range args {
		if arg == na.Name {
			return na.Name, true
		}
	}
	return "", false
}

func Register(name string, a Action) {
	actions.Register(name, a)
}

func Actions() ActionMap{
	return actions
}

type Action interface {
	flags.HasFlags
	Process(resource *resource.Resource)
}
