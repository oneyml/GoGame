package module

import (
	send "GoGame/core/libs/sendModule"
	"GoGame/core/libs/sessions"
	"GoGame/core/protos/gameProto"
	"GoGame/servives/server/module/chatroom"
	"log"
	"github.com/golang/protobuf/proto"
)

func ChatJoin(session *sessions.FrontSession, msgBody proto.Message) {
	log.Println("ChatJoin ok")
	chatroom.JoinRoom(session.RId)
	nmsgBody := &gameProto.ClientChatJoinS2C{}
	send.SendMsgToClient(session, nmsgBody)
	chatroom.SendHisMsg(session)
}

func Chat(session *sessions.FrontSession, msgBody proto.Message) {
	log.Printf("player:%s Chat ok", session.RName)
	data := msgBody.(*gameProto.ClientChatC2S)
	text := chatroom.Filter(data.GetMsg())
	chatroom.RegMsg(session, text)

	msgBody1 := &gameProto.ClientChatS2C{
		SpeakerId:   proto.Uint64(session.RId),
		SpeakerName: proto.String(session.RName),
		SpeakerText: proto.String(text),
	}
	chatroom.Broadcast(session.RId, msgBody1)
}

func Popular(session *sessions.FrontSession, msgBody proto.Message) {
	data := msgBody.(*gameProto.ClientChatPopularC2S)
	log.Printf("player:%s n %d", session.RName, data.GetN())
	word := chatroom.GetPopular(session.RId, data.GetN())
	msgBody1 := &gameProto.ClientChatPopularS2C{
		World: proto.String(word),
	}
	send.SendMsgToClient(session, msgBody1)
}

