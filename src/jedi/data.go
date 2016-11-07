package jedi

// QINIU_JEDI_HOST API域名
const QINIU_JEDI_HOST = "http://jedi.qiniuapi.com"

// TransferConf 预设配置
type TransferConf struct {
	Name    string    `json:"name"`
	Enabled bool      `json:"enabled"`
	Format  string    `json:"format"`
	Video   VideoInfo `json:"video"`
	Audio   AudioInfo `json:"audio"`
}

// ***视频信息 Start***
type videoItems struct {
	Total  int         `json:"total"`
	Count  int         `json:"count"`
	Cursor string      `json:"cursor"`
	Items  []VideoItem `json:"items"`
}

type VideoItem struct {
	Key              string        `json:"key"`
	Name             string        `json:"name"`
	Tags             []string      `json:"tags"`
	Description      string        `json:"description"`
	Size             int           `json:"size"`
	CreationTime     string        `json:"creation_time"`
	ModificationTime string        `json:"modification_time"`
	Status           string        `json:"status"`
	AvInfo           AvInformation `json:"avinfo"`
	TransCoding      []TransCode   `json:"transcoding"`
	ThumbNails       ThumbnailInfo `json:"thumbnails"`
	UserInfo         interface{}   `json:"user_info"`
}

type AvInformation struct {
	Format AvFormat  `json:"format"`
	Video  VideoInfo `json:"video"`
	Audio  AudioInfo `json:"audio"`
}

// VideoInfo 视频参数
type VideoInfo struct {
	BitRate   string `json:"bit_rate"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	FrameRate string `json:"frame_rate"`
	CodecName string `json:"codec_name"`
}

// AudioInfo 音频参数
type AudioInfo struct {
	BitRate    string `json:"bit_rate"`
	Quality    int    `json:"quality"`
	SampleRate string `json:"sample_rate"`
	CodecName  string `json:"codec_name"`
}

type AvFormat struct {
	Duration       string `json:"duration"`
	FormatLongName string `json:"format_long_name"`
}

type TransCode struct {
	TransSetID    string `json:"tansset_id"`
	TransSetName  string `json:"transset_name"`
	TransConfID   string `json:"transconf_id"`
	TransConfName string `json:"transconf_name"`
	Url           string `json:"url"`
	Status        string `json:"status"`
	Format        string `json:"format"`
}

type ThumbnailInfo struct {
	Urls   []string `json:"urls"`
	Active int      `json:"active"`
}

// ***视频信息 End ***

type DeleteInfo struct {
	Errors bool         `json:"errors"`
	Items  []DeleteItem `json:"items"`
}
type DeleteItem struct {
	Key    string `json:"key"`
	Status int    `json:"status"`
}

type WaterMarkConf struct {
	Type      string `json:"type"`
	Enabled   bool   `json:"enabled"`
	Gravity   string `json:"gravity"`
	OffsetX   int    `json:"offset_x"`
	OffsetY   int    `json:"offset_y"`
	Image     string `json:"image"`
	Text      string `json:"text"`
	TextFont  string `json:"text_font"`
	TextColor string `json:"text_color"`
	TextSize  int    `json:"text_size"`
}
type WaterMarkGroup struct {
	Count int             `json:"count"`
	Items []WaterMarkConf `json:"items"`
}
type WaterMarkCrt struct {
	ID string `json:"id"`
}
