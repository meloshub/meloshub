# MelosHub

MelosHub是一个跨平台音乐资源搜索的中间层。它提供一个统一的抽象接口标准，并使用适配器模式工作。

命名由来：melos这个词源自古希腊语“Melos”（μέλος），意为旋律，音律，曲调，歌曲。hub就是枢纽，集线器。~~所以你可以称本项目为“旋律集线器”。~~

> [!NOTE]
> 此项目还在积极开发中，项目中的函数定义、项目结构、用法可能会不断变化。

## 使用

安装：

```bash
go get https://github.com/meloshub/meloshub
```

运行：

```
go run main.go
```

### 初始化适配器

初始化日志模块：

```go

import "github.com/meloshub/meloshub/logging"

logging.Init(logging.Config{
		Level:     "info",
		Format:    "consle",
		AddSource: true,
	})
slog.Info("Meloshub is initializing..")
```

获取所有可用适配器：

```go
import "github.com/meloshub/meloshub/adapter"

slog.Info("Getting available adapters")
adapters := adapter.GetAll()
if len(adapters) == 0 {
    slog.Warn("Adapters list is empty")
}
for _, v := range adapters {
    println(v.Id())
}
```

### 适配器基础API

使用指定适配器进行音乐搜索：

```go
qqmusic, ok := adapter.Get("qqmusic")
if !ok {
		slog.Error("adapter 'example' is not existed")
	}
qqmusic.Search("夜的第七章", adapter.SearchOptions{
    Page:  1,
    Limit: 10,
})
```

根据平台的歌曲id播放歌曲：

```go
playUrl, err := qqmusic.PlayURL("004Ng8xu20eirf")
if err != nil {
	slog.Error(err.Error())
	return
}
slog.Info(fmt.Sprint("Got play url: ", playUrl))
```

根据平台的歌曲id获取歌词：

```go
lyrics, err := qqmusic.Lyrics(songList[0].ID)
if err != nil {
	slog.Error(err.Error())
	return
}
slog.Info(fmt.Sprintln("Lyrics: \n", lyrics))
```

## 编写适配器（开发者）

适配器的代码与相关问题应该被提交到[适配器仓库](https://github.com/meloshub/meloshub-adapters)。如果你有新的适配想法，欢迎提交issue，有能力的开发者可以提交PR至此仓库。

