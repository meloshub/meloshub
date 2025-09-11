package example

import (
	"fmt"

	"github.com/meloshub/meloshub/adapter"
	"github.com/meloshub/meloshub/model"
)

type ExampleAdapter struct {
	adapter.Base
}

// 适配器将在导入时被注册
func init() {
	if err := adapter.Register(New()); err != nil {
		panic(fmt.Errorf("failed to register adapter: %w", err))
	}
}

func New() *ExampleAdapter {
	a := &ExampleAdapter{}
	metadata := adapter.Metadata{
		Id:          "example",
		Title:       "Example Adapter",
		Type:        adapter.TypeCommunity,
		Version:     "1.0.0",
		Author:      "meloshub",
		Description: "An example adapter for meloshub development",
	}
	a.Init(metadata)
	return a
}

func (a *ExampleAdapter) SearchSong(keyword string, options adapter.SearchOptions) ([]model.Song, error) {
	return []model.Song{}, nil
}

func (a *ExampleAdapter) GetSongByID(id string) (*model.Song, error) {
	return &model.Song{
		ID: id,
	}, nil
}

func (a *ExampleAdapter) GetLyricsByID(id string) (string, error) {
	return "", nil
}

func (a *ExampleAdapter) GetAlbumSongsByID(id string) ([]model.Song, error) {
	return []model.Song{}, nil
}
