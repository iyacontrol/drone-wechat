package wechat

import (
	"github.com/iyacontrol/drone-wechat/wechat/context"
	"github.com/iyacontrol/drone-wechat/wechat/util"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	CorpID         string
	CorpSecret     string
	Token          string
	EncodingAESKey string
}

// NewWechat init
func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.CorpID = cfg.CorpID
	context.CorpSecret = cfg.CorpSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
}

// SendText 发送文本消息
func (wc *Wechat) SendText(text string) error {
	reply, err := util.PostJSON("")
}
