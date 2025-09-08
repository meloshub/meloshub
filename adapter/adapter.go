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

// Adapter 是所有音乐平台适配器必须实现的接口。
// 它定义了从不同音乐源获取数据的标准化方法。
type Adapter interface {
	// Platform 返回适配器的平台标识符
	// 这个标识符应该全局唯一且为小写字符串。
	Platform() string

	// SearchSong 根据关键词搜索歌曲。
	// 返回一个歌曲切片和可能发生的错误。
	SearchSong(keyword string, options SearchOptions) ([]model.Song, error)

	// GetSongByID 根据歌曲ID获取歌曲的详细信息。
	GetSongByID(id string) (*model.Song, error)

	// GetAlbumSongsByID 根据专辑ID获取该专辑下的所有歌曲。
	GetAlbumSongsByID(id string) ([]model.Song, error)
}

type Base struct {
	PlatformName string           //平台唯一标识
	Session      *network.Session //会话实例
	Logger       *slog.Logger     //日志上下文实例
	Config       map[string]any   //适配器配置项
}

// Init 初始化 BaseAdapter
func (b *Base) Init(platform string) {
	b.PlatformName = platform
	b.Session = network.NewSession()
	b.Logger = logging.Get().With("adapter", platform)
	b.Config = make(map[string]any)
}

// Platform 获取适配器对应平台的唯一标识符
func (b *Base) Platform() string {
	return b.PlatformName
}
