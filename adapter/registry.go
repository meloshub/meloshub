package adapter

import (
	"errors"
	"fmt"
	"log"
	"maps"
)

var providers = make(map[string]Adapter)

var (
	ErrAdapterIsNil           = errors.New("registry: adapter instance cannot be nil")
	ErrAdapterIdIsEmpty       = errors.New("registry: adapter id cannot be empty")
	ErrAdapterMetaDataIsEmpty = errors.New("registry: adapter metadata cannot be empty")
	ErrAdapterIdExists        = errors.New("registry: adapter id already exists")
)

// Register 注册一个适配器实例到全局适配器注册表中
// 如果适配器实例为 nil，平台名称为空，或者平台名称已存在，则返回错误
func Register(a Adapter) error {
	if a == nil {
		return ErrAdapterIsNil
	}

	meta := a.Metadata()
	if (meta == Metadata{}) {
		return ErrAdapterMetaDataIsEmpty
	}

	if meta.Id == ""{
		return ErrAdapterIdIsEmpty
	}

	if _, exists := providers[meta.Id]; exists {
		return fmt.Errorf("%w: %s", ErrAdapterIdExists, meta.Id)
	}

	providers[meta.Id] = a
	log.Printf("Adapter '%s' registered successfully.", meta.Id)
	return nil
}

// Get 从注册表中获取指定平台的适配器实例
func Get(id string) (Adapter, bool) {
	p, ok := providers[id]
	return p, ok
}

// GetAll 返回所有已注册的适配器实例的副本
func GetAll() map[string]Adapter {
	clone := make(map[string]Adapter)
	maps.Copy(clone, providers)
	return clone
}
