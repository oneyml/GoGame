package module

import (
	send "GoGame/core/libs/sendModule"
	"GoGame/core/libs/sessions"
	"GoGame/core/protos/gameProto"
	"github.com/golang/protobuf/proto"
	"log"
	"strconv"
	"time"
)

func con(d int64, unit string) string {
	str := ""
	if d > 10 {
		str = " " + strconv.FormatInt(d, 10)+unit
	} else if d > 0 {
		str = " 0" + strconv.FormatInt(d, 10)+unit
	} else {
		str = " 00"+unit
	}
	return str
}

func GetStats(last int64) string {
	str := ""
	now := time.Now().Unix()
	diff := now - last

	day    := diff/(60*60*24)
	str += con(day, "d")
	diff -= day*60*60*24

	hour   := diff/(60*60)
	str += con(hour, "h")
	diff -= hour*60*60

	minute := diff/60
	str += con(minute, "m")

	second := diff%60
	str += con(second, "s")

	return str
}

func Stats(session *sessions.FrontSession, msgBody proto.Message) {
	data := msgBody.(*gameProto.ClientChatStatsC2S)
	log.Printf("player:%s Stats %s\n", session.RName, data.GetRname())
	name := data.GetRname()
	var se *sessions.FrontSession
	sessions.FetchFrontSession(func(clientSession *sessions.FrontSession) {
		if clientSession.RName == name {
			se = clientSession
		}
	})
	msgBody1 := &gameProto.ClientChatStatsS2C{
		Rname: proto.String(data.GetRname()),
		Time : proto.String(""),
	}
	if se != nil {
		msgBody1.Time = proto.String(GetStats(se.Time))
	}
	send.SendMsgToClient(session, msgBody1)
}