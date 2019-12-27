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

func (am *ActionMap) SearchAndConsumeArgument(args []string) (Action, []string, bool) {
	for _, namedAction := range *am {
		_, cargs, ok := namedAction.searchAndConsumeArgument(args)
		if ok {
			return namedAction.Action, cargs, ok
		}
	}
	return nil, args, false
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

func (na *NamedAction) searchAndConsumeArgument(args []string) (string, []string, bool) {
	result_name := ""
	result_ok := false
	consumed_args := []string{}
	for _, arg := range args {
		if arg == na.Name {
			result_name = na.Name
			result_ok   = true
			continue
		}
		consumed_args = append(consumed_args, arg)
	}
	return result_name, consumed_args, result_ok
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
