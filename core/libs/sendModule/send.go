package send

import (
	"GoGame/core/protos"
	"github.com/golang/protobuf/proto"
)

type sends interface {
	Send([]byte) error
}

func SendMsgToClient(session sends, sendMsg proto.Message) {
	session.Send(protos.MarshalProtoMsg(sendMsg))
}
