package server

import (
	"context"
	"strconv"

	"github.com/iyacontrol/drone-wechat/wechat/util"
	"github.com/iyacontrol/drone-wechat/wechat/context"

//Server struct
type Server struct {
	*context.Context

	openID string

	responseRawXMLMsg []byte
	responseMsg       interface{}

	isSafeMode bool
	random     []byte
	nonce      string
	timestamp  int64
}

//NewServer init
func NewServer(context *context.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

//GetOpenID return openID
func (srv *Server) GetOpenID() string {
	return srv.openID
}

//Send 将自定义的消息发送
func (srv *Server) Send() (err error) {
	replyMsg := srv.responseMsg
	if srv.isSafeMode {
		//安全模式下对消息进行加密
		var encryptedMsg []byte
		encryptedMsg, err = util.EncryptMsg(srv.random, srv.responseRawXMLMsg, srv.AppID, srv.EncodingAESKey)
		if err != nil {
			return
		}
		//TODO 如果获取不到timestamp nonce 则自己生成
		timestamp := srv.timestamp
		timestampStr := strconv.FormatInt(timestamp, 10)
		msgSignature := util.Signature(srv.Token, timestampStr, srv.nonce, string(encryptedMsg))
		replyMsg = message.ResponseEncryptedXMLMsg{
			EncryptedMsg: string(encryptedMsg),
			MsgSignature: msgSignature,
			Timestamp:    timestamp,
			Nonce:        srv.nonce,
		}
	}
	if replyMsg != nil {
		srv.XML(replyMsg)
	}
	return
}
