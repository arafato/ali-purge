package resources

import (
	"fmt"
	"sync"

	"github.com/arafato/ali-purge/cmd"
)

type ResourceListerFunc func(config *cmd.AlicloudConfig) ([]cmd.Resource, error)

type ResourceManager struct {
	resources map[string]ResourceListerFunc
}

var smInstance *ResourceManager
var onceSm sync.Once

func GetServiceManager() *ResourceManager {
	onceSm.Do(func() {
		smInstance = &ResourceManager{}
		smInstance.resources = make(map[string]ResourceListerFunc)
	})
	return smInstance
}

func (rm *ResourceManager) register(name string, lister ResourceListerFunc) {
	_, exists := rm.resources[name]
	if exists {
		panic(fmt.Sprintf("A resource with the name %s already exists", name))
	}

	rm.resources[name] = lister
}
