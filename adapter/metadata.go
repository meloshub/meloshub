package adapter

// AdapterType 适配器类型
type AdapterType string

const (
	// TypeOfficial 由 meloshub 官方维护的适配器
	TypeOfficial AdapterType = "official"
	// TypeCommunity 由社区开发者贡献和维护的适配器
	TypeCommunity AdapterType = "community"
)

// Metadata 适配器元数据信息
type Metadata struct {
	// Id 适配器的唯一标识符
	// 必须是小写字母或数字
	Id string `json:"id" yaml:"id"`

	// Title 适配器的显示名称
	Title string `json:"title" yaml:"title"`

	// Type 适配器类型：官方适配器或社区适配器
	Type AdapterType `json:"type" yaml:"type"`

	// Version 适配器的版本号，必须遵循语义化版本
	Version string `json:"version" yaml:"version"`

	// Author 适配器作者的名称或id
	Author string `json:"author" yaml:"author"`

	// Description 适配器描述
	Description string `json:"description" yaml:"description"`
}
