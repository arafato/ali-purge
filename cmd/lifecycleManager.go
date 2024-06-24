package cmd

import (
	"sync"
)

type Resource interface {
	Remove() error
}

type ItemState int

const (
	ItemStateNew ItemState = iota
	ItemStatePending
	ItemStateFailed
	ItemStateFinished
)

type Item struct {
	Resource Resource
	State    ItemState
}

type LifecycleManager struct {
	items []*Item
}

var rmInstance *LifecycleManager
var onceRm sync.Once

func GetResourceManager() *LifecycleManager {
	onceRm.Do(func() {
		rmInstance = &LifecycleManager{}
		rmInstance.items = make([]*Item, 0)
	})
	return rmInstance
}
