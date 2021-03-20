package module

import (
	"GoGame/core/libs/sessions"
	"github.com/golang/protobuf/proto"
	"log"
)

func Ping(session *sessions.FrontSession, msgBody proto.Message) {
	log.Println("ping", session.RId, session.RName)
}
