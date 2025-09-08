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
	a.Init("example")
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

func (a *ExampleAdapter) GetAlbumSongsByID(id string) ([]model.Song, error) {
	return []model.Song{}, nil
}
