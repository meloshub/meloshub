package adapter

import (
	"log/slog"

	"github.com/meloshub/meloshub/logging"
	"github.com/meloshub/meloshub/model"
	"github.com/meloshub/meloshub/network"
)

// SearchOptions 定义了搜索时可以传入的额外参数
type SearchOptions struct {
	Page  int
	Limit int
}

// Adapter 是所有音乐平台适配器必须实现的接口
type Adapter interface {
	// Metadata 返回适配器的元数据信息
	Metadata() Metadata

	// Id 获取适配器唯一标识符
	Id() string

	// Search 根据关键词搜索歌曲。
	// 返回一个歌曲切片和可能发生的错误。
	Search(keyword string, options SearchOptions) ([]model.Song, error)

	// PlayURL 获取歌曲播放链接。
	PlayURL(id string) (string, error)

	// Lyrics 获取歌词信息。
	Lyrics(id string) (string, error)

	// AlbumDetail 根据专辑ID获取该专辑的详细信息。
	AlbumDetail(id string) (model.Album, error)
}

type Base struct {
	id       string           //适配器唯一标识
	metadata Metadata         //适配器元数据
	Session  *network.Session //会话实例
	Logger   *slog.Logger     //日志上下文实例
	Config   map[string]any   //适配器配置项
}

// Init 初始化 BaseAdapter
func (b *Base) Init(meta Metadata) {
	b.id = meta.Id
	b.metadata = meta
	b.Session = network.NewSession()
	b.Logger = logging.Get().With("adapter", meta.Id)
	b.Config = make(map[string]any)
}

// Id 获取适配器对应平台的唯一标识符
func (b *Base) Id() string {
	return b.id
}

func (b *Base) Metadata() Metadata {
	return b.metadata
}
