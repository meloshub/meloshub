package main

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/meloshub/meloshub/adapter"
	"github.com/meloshub/meloshub/logging"

	_ "github.com/meloshub/meloshub-adapters/all" //加载仓库所有可用适配器
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
		fmt.Printf("adapter：%s is available", v.Id())
	}

	// 获取指定的适配器并使用
	qqmusic, ok := adapter.Get("qqmusic")
	if !ok {
		slog.Error("adapter 'qqmusic' is not existed")
	}

	songList, err := qqmusic.Search("夜的第七章", adapter.SearchOptions{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		slog.Error(err.Error())
		return
	}

	jsonBytes, err := json.MarshalIndent(songList, "", "  ")
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
