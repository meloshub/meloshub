# MelosHub

MelosHub是一个跨平台音乐资源搜索的中间层。它提供一个统一的抽象接口标准，并使用适配器模式工作。

## 使用

安装：

```bash
go get https://github.com/meloshub/meloshub
```

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
    println(v.Platform())
}
```

使用指定适配器进行音乐搜索：

```go
exampleAdapter, ok := adapter.Get("example")
if !ok {
		slog.Error("adapter 'example' is not existed")
	}
exampleAdapter.SearchSong("夜的第七章", adapter.SearchOptions{
    Page:  1,
    Limit: 10,
})
```

## 编写适配器（开发者）

适配器的代码与相关问题应该被提交到[适配器仓库](https://github.com/meloshub/meloshub-adapters)。如果你有新的适配想法，欢迎提交issue，有能力的开发者可以提交PR至此仓库。

### 1. 定义适配器

定义一个新的适配器结构体，并继承adapter.Base：

```go
type ExampleAdapter struct {
	adapter.Base
}
```

### 2. 实现接口

需要实现adapter.Adapter中定义的所有方法：

```go
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
```

### 3. 注册适配器

在构造函数中使用Init方法初始化适配器，在init中使用adapter.Register接收构造函数：

```go
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
```

### 4. 导入适配器

在主程序入口中导入适配器：

```go
import _ "github.com/meloshub/meloshub-adapters/example"
```

