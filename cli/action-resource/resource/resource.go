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

func (rm *ResourceMap) Search(args []string) (Resource, bool) {
	for _, namedResource := range *rm {
		_, ok := namedResource.search(args)
		if ok {
			return namedResource.Resource, ok
		}
	}
	return nil, false
}

type NamedResource struct {
	Name string
	Resource Resource
}

func (nr *NamedResource) search(args []string) (string, bool) {
	for _, arg := range args {
		if arg == nr.Name {
			return nr.Name, true
		}
	}
	return "", false
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
