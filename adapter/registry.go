package adapter

import (
	"errors"
	"fmt"
	"log"
)

var providers = make(map[string]Adapter)

var (
	ErrAdapterIsNil    = errors.New("registry: adapter instance cannot be nil")
	ErrPlatformIsEmpty = errors.New("registry: adapter platform name cannot be empty")
	ErrPlatformExists  = errors.New("registry: adapter for this platform already exists")
)

// Register 注册一个适配器实例到全局适配器注册表中
// 如果适配器实例为 nil，平台名称为空，或者平台名称已存在，则返回错误
func Register(a Adapter) error {
	if a == nil {
		return ErrAdapterIsNil
	}

	platform := a.Platform()
	if platform == "" {
		return ErrPlatformIsEmpty
	}

	if _, exists := providers[platform]; exists {
		return fmt.Errorf("%w: %s", ErrPlatformExists, platform)
	}

	providers[platform] = a
	log.Printf("Adapter '%s' registered successfully.", platform)
	return nil
}

// Get 从注册表中获取指定平台的适配器实例
func Get(platform string) (Adapter, bool) {
    p, ok := providers[platform]
    return p, ok
}

// GetAll 返回所有已注册的适配器实例的副本
func GetAll() map[string]Adapter {
    clone := make(map[string]Adapter)
    for k, v := range providers {
        clone[k] = v
    }
    return clone
}
