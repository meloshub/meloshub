package model

// Artist 艺术家信息
type Artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Album 专辑信息
type Album struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	CoverURL string `json:"cover_url"`
}

// Audio 音频文件信息
type AudioFile struct {
	Md5     string `json:"md5"`
	Quality int    `json:"quality"`
	Format  string `json:"format"`
	FileSize    int64  `json:"size"`
	URL     string `json:"url"`
}

// Song 统一的歌曲抽象模型
type Song struct {
	ID         string   `json:"id"`               // 歌曲在源平台的唯一ID
	Source     string   `json:"source"`           // 源平台标识, e.g., "netease", "qqmusic", "kuwo"
	Title      string   `json:"title"`            // 歌曲标题
	Artists    []Artist `json:"artists"`          // 艺术家列表
	Album      Album    `json:"album"`            // 所属专辑
	AudioURL   string   `json:"audio_url"`        // 获取到的音频播放链接
	DurationMs int      `json:"duration_ms"`      // 歌曲时长（毫秒）
	AudioFiles []AudioFile `json:"audio_files,omitempty"` // 可选的多质量音频文件列表，如果有的话
	Playable   bool     `json:"playable"`         // 当前是否可播放（可能因地区限制等不可播放）
	IsVIP      bool     `json:"is_vip"`           // 在源平台是否为VIP歌曲
	Lyrics     string   `json:"lyrics,omitempty"` // 歌词文本（可选）
}
