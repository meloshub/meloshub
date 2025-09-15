package main

import (
	"log/slog"

	"github.com/meloshub/meloshub/adapter"
	"github.com/meloshub/meloshub/logging"
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
	_, ok := adapter.Get("example")
	if !ok {
		slog.Error("adapter 'example' is not existed")
	}
}
