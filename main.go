package main

import (
	"log/slog"

	"github.com/meloshub/meloshub/adapter"
	"github.com/meloshub/meloshub/logging"

	_ "github.com/meloshub/meloshub/adapter/example"
)

func main() {
	logging.Init(logging.Config{
		Level:     "info",
		Format:    "consle",
		AddSource: true,
	})
	slog.Info("Meloshub is initializing..")
	slog.Info("Getting available adapters")
	adapters := adapter.GetAll()
	if len(adapters) == 0 {
		slog.Warn("Adapters list is empty")
	}
	for _, v := range adapters {
		println(v.Id())
	}
	exampleAdapter, ok := adapter.Get("example")
	if !ok {
		slog.Error("adapter 'example' is not existed")
	}
	exampleAdapter.SearchSong("夜的第七章", adapter.SearchOptions{
		Page:  1,
		Limit: 10,
	})
}
