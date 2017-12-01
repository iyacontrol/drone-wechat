package message

import "github.com/iyacontrol/drone-wechat/wechat/util"

// MsgType 基本消息类型
type MsgType string

const (
	//MsgTypeText 表示文本消息
	MsgTypeText MsgType = "text"
	//MsgTypeImage 表示图片消息
	MsgTypeImage = "image"
	//MsgTypeVoice 表示语音消息
	MsgTypeVoice = "voice"
	//MsgTypeVideo 表示视频消息
	MsgTypeVideo = "video"
	//MsgTypeFile 表示文件
	MsgTypeFile = "file"
	//MsgTypeTextCard 表示文本卡片
	MsgTypeTextCard = "textcard"
	//MsgTypeMpNews 表示mp图文消息[限回复]
	MsgTypeMpNews = "mpnews"
	//MsgTypeNews 表示图文消息[限回复]
	MsgTypeNews = "news"
)

// ResponseJsonMsg 需要返回的消息体
type ResponseJsonMsg struct {
	util.CommonError
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

// CommonContent 消息中通用的结构
type CommonContent struct {
	ToUser  string  `json:"touser"`
	ToParty string  `json:"toparty"`
	ToTag   string  `json:"totag"`
	AgentID int     `json:"agentid"`
	MsgType MsgType `json:"msgtype"`
}

//SetToUser set ToUser
func (msg *CommonContent) SetToUser(toUser string) {
	msg.ToUser = toUser
}

//SetToParty set ToParty
func (msg *CommonContent) SetToParty(toParty string) {
	msg.ToParty = toParty
}

//SetToTag set ToTag
func (msg *CommonContent) SetToTag(toTag string) {
	msg.ToTag = toTag
}

//SetAgentID set AgentID
func (msg *CommonContent) SetAgentID(agentID int) {
	msg.AgentID = agentID
}

//SetMsgType set MsgType
func (msg *CommonContent) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}
