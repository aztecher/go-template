package resource

import (
	"github.com/go-template/cli/action-resource/flags"
)

var resources = ResourceMap{}

type ResourceMap []NamedResource

func (rm *ResourceMap) Register(name string, r Resource) {
	nr := NamedResource{ Name: name, Resource: r}
	*rm = append(*rm, nr)
}

func (rm *ResourceMap) SearchAndConsumeArgument(args []string) (Resource, []string, bool) {
	for _, namedResource := range *rm {
		_, cargs, ok := namedResource.searchAndConsumeArgument(args)
		if ok {
			return namedResource.Resource, cargs, ok
		}
	}
	return nil, args, false
}


type NamedResource struct {
	Name string
	Resource Resource
}

func (nr *NamedResource) searchAndConsumeArgument(args []string) (string, []string, bool) {
	result_name := ""
	result_ok := false
	consumed_args := []string{}
	for _, arg := range args {
		if arg == nr.Name {
			result_name = nr.Name
			result_ok   = true
			continue
		}
		consumed_args = append(consumed_args, arg)
	}
	return result_name, consumed_args, result_ok
}

func Register(name string, r Resource) {
	resources.Register(name, r)
}

func Resources() ResourceMap {
	return resources
}

type Resource interface {
	flags.HasFlags
	Process()
}
