package module

import (
	send "GoGame/core/libs/sendModule"
	"GoGame/core/libs/sessions"
	"GoGame/core/protos/gameProto"
	"github.com/golang/protobuf/proto"
	"log"
)

func deleteRepeatSession(rid uint64) {
	session := sessions.GetFrontSession(rid)
	if session != nil {
		session.Close()
	}
}

func checkUniqName(name string) bool  {
	find := true
	sessions.FetchFrontSession(func(clientSession *sessions.FrontSession) {
		if clientSession.RName == name {
			find = false
		}
	})
	return find
}

func Login(session *sessions.FrontSession, msgBody proto.Message) {
	data := msgBody.(*gameProto.ClientLoginC2S)
	rid  := data.GetRId()
	rname := data.GetRName()
	msgBody1 := &gameProto.ClientLoginS2C{
		RId:   proto.Uint64(rid),
		Token: proto.String("token"),
		Error: proto.Uint64(0),
	}
	if checkUniqName(rname) != true {
		msgBody1.Error = proto.Uint64(1)
	}else{
		session.RName = rname
		session.RId = rid
		log.Printf("login ok:%d, %s", rid, session.RName)
		deleteRepeatSession(rid)
		sessions.AddFrontSession(session)
	}
	send.SendMsgToClient(session, msgBody1)
}
