package resource

import (
	"fmt"
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

func (rm *ResourceMap) List() string {
	result := ""
	for _, ns := range *rm {
		nsStr := fmt.Sprintf("  %s     %s\n", ns.Name, ns.Name)
		result = result + nsStr
	}
	return result
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
	// ここが処理の表現力に当たる
	// このinterfaceがAction interfaceに渡され,
	// 各action interfaceの実装の表現力の上で、これを利用することになる
	// action-resourceで関係の保持が必要なので、
	// その分処理を多重化するためにrequirementが増える
	// 緩和法は思い浮かばない、interfaceで切り分けして統合するくらいか
	// TODO: relationを切り分け束ねるinterfaceを作っていくのがいいのかも
	flags.HasFlags
	Process()
	ShowProcess() // こういうのが増える(action-resource relation)
}
