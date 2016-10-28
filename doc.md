PACKAGE DOCUMENTATION

package jedi
    import "/Users/tianyang/Documents/go_wrkspace/jedi-golang-sdk/src-1.7/jedi/"


CONSTANTS

const QINIU_JEDI_HOST = "http://jedi.qiniuapi.com"
    QINIU_JEDI_HOST API域名

VARIABLES

var ACCESS_KEY string
    ACCESS_KEY qiniu access key

var SECRET_KEY string
    SECRET_KEY qiniu secret key

FUNCTIONS

func BatchdeleteVideos(hub string, videoKeys []string) string
    BatchdeleteVideos 批量删除 videoKeys 视频列表

func CreateConf(hub, name string) string
    CreateConf 创建转码配置

func CreatePreset(hub, transID string, config TransferConf) string
    CreatePreset 创建转码预设

func DeleteConf(hub, transID string) string
    DeleteConf 删除转码配置

func DeletePreset(hub, transID, transSetID string) string
    DeletePreset 删除转码预设

func DeleteVideo(hub, key string) string
    DeleteVideo 删除视频

func EnablePreset(hub, transID, transSetID, enable string) []byte
    EnablePreset 启用禁用转码预设

func GetPreset(hub, transID, transSetID string) string
    GetPreset 获取转码预设

func GetUpToken(hub, videoType string, deadline int) string
    GetUpToken 获取uptoken Hub 点播空间名 deadline 有效时间 aksk 用户密钥

func GetVideoInfo(hub, key string) string
    GetVideoInfo 获取视频信息 <Hub>:点播空间名 <key>: 对video key(对应源bucket key)做
    urlsafe的base64编码

func GetVideoList(hub, cursor string, count int) string
    GetVideoList 获取视频信息列表

func ListTransConf(hub string) string
    ListTransConf 获取所有转码配置

func RequestWithBody(method, url string, body interface{}) []byte
    RequestWithBody 带body对api发出请求并且返回response body

func RequestWithoutBody(method, url string) []byte
    RequestWithoutBody 对api发出请求并且返回response body

func SetMac(accessKey, secretKey string)
    SetMac 初始化ak,sk

func SetVideoImage(hub, key string, active int) string
    SetVideoImage 设置封面 active 封面索引

func Sign(req *http.Request) (token string)
    Sign return the QiniuToken for Authorization Every request to jedi API
    must Add a Header Authorization which value is tis token

func UpdateVideoInfo(hub, key, name, description string, tags []string) string
    UpdateVideoInfo 更新视频信息

func UpgradePreset(hub, transID, transSetID string, config TransferConf) string
    UpgradePreset 更新转码预设

func UploadVideoFile(hub, videoType, filePath, key string, deadline int) error
    UploadVideoFile 上传文件 Hub 点播空间名 videoType video类型 filePath 文件目录 key 保存文件名
    deadline token有效时间

TYPES

type AudioInfo struct {
    BitRate    string `json:"bit_rate"`
    Quality    int    `json:"quality"`
    SampleRate string `json:"sample_rate"`
    CodecName  string `json:"codec_name"`
}
    AudioInfo 音频参数

type TransferConf struct {
    Name    string    `json:"name"`
    Enabled bool      `json:"enabled"`
    Format  string    `json:"format"`
    Video   VideoInfo `json:"video"`
    Audio   AudioInfo `json:"audio"`
}
    TransferConf 预设配置

type UploadTokenBody struct {
    Hub      string `json:"hub"`
    DeadLine int    `json:"deadline"`
    Type     string `json:"type"`
}
    UploadTokenBody 上传参数

type VideoInfo struct {
    BitRate   string `json:"bit_rate"`
    Width     int    `json:"width"`
    Height    int    `json:"height"`
    FrameRate string `json:"frame_rate"`
    CodecName string `json:"codec_name"`
}
    VideoInfo 视频参数


