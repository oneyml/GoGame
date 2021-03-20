package messages

import (
	"GoGame/core/libs/sessions"
	"GoGame/core/protos"
	"github.com/golang/protobuf/proto"
	"log"
)

var (
	MapHandle = make(map[uint16]ipcServerMsgHandle)
)

type ipcServerMsgHandle func(session *sessions.FrontSession, msgBody proto.Message)

func GetIpcServerHandle(msgId uint16) ipcServerMsgHandle {
	handle := MapHandle[msgId]
	if handle == nil {
		log.Println("unregister msgId:", msgId)
	}
	return handle
}

func RegisterIpcServerHandle(msgID uint16, handle ipcServerMsgHandle) {
	MapHandle[msgID] = handle
}

func MsgHandle(session *sessions.FrontSession, msgBody []byte) {
	msgProto := protos.UnMarshalProtoMsg(msgBody)
	handle := GetIpcServerHandle(msgProto.MsgId)
	if handle != nil {
		handle(session, msgProto.MsgBody)
	}

}
