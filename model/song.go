package model

// Singer 歌手信息
type Singer struct {
	ID   string `json:"id"`   //歌手id
	Name string `json:"name"` //歌手名称
}

// Album 专辑信息
type Album struct {
	ID              string `json:"id"`               //专辑id
	Name            string `json:"name"`             //专辑名称
	Description     string `json:"description"`      //专辑描述
	PublicTimestamp int64  `json:"public_timestamp"` //发布时间戳
	CoverURL        string `json:"cover_url"`        //封面链接
	SongList        []Song `json:"song_list"`        //歌曲信息列表
}

// AudioFile 音频文件信息
type AudioFile struct {
	Name     string `json:"name"`    // 音频名称
	Format   string `json:"format"`  // 音频格式
	Quality  int    `json:"quality"` // 音频质量
	FileSize int64  `json:"size"`    // 文件大小
	Bytes    string `json:"bytes"`   // 音频的二进制数据
	Md5      string `json:"md5"`     // 二进制数据的32位md5编码
}

// Song 统一的歌曲信息抽象模型
type Song struct {
	ID        string   `json:"id"`         // 歌曲在源平台的唯一ID
	Source    string   `json:"source"`     // 源平台标识
	Title     string   `json:"title"`      // 歌曲标题
	Singers   []Singer `json:"singers"`    // 歌手列表
	AlbumId   string   `json:"album_id"`   // 所属专辑id
	AlbumName string   `json:"album_name"` // 所属专辑名称
	Playable  bool     `json:"playable"`   // 当前是否可播放
}
