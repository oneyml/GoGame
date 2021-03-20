package chatroom

import (
	"GoGame/core/config"
	"GoGame/core/libs/guid"
	send "GoGame/core/libs/sendModule"
	"GoGame/core/libs/sessions"
	"GoGame/core/protos/gameProto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"sync"
	"time"
)

var (
	maxMsgMap = int(config.GetChatHistory())
	room      = rooms{}
	guid_        = guid.NewGuid(1)

	ridJoinMap   = make(map[uint64]*ChatRooms)
	rnameJoinMap = make(map[string]uint64)
	roomMap      = make(map[uint64]*ChatRooms)
	romMutex   sync.RWMutex
)

type Msg struct {
	Rid   uint64
	Rname string
	Text  string
	Time  string
}

type rooms struct {
}

type joinPlayers struct {
	joinTime time.Time
	rid      uint64
	rname    string
}

type ChatRooms struct {
	roomId   		uint64
	roomName 		string
	joinPlayerMap 	map[uint64]joinPlayers
	time 		    int64
	msgMap          []Msg
	popular         *populars
}

func timeStr() string {
	time.Now().String()
	t1 := time.Now().Year()
	t2 := time.Now().Month()
	t3 := time.Now().Day()
	t4 := time.Now().Hour()
	t5 := time.Now().Minute()
	t6 := time.Now().Second()
	t7 := time.Now().Nanosecond()
	currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local).String()
	return currentTimeData
}

func packMsg(session *sessions.FrontSession, text string) Msg {
	return Msg{
		Rid:   session.RId,
		Rname: session.RName,
		Text:  text,
		Time:  time.Now().String(),
	}
}

func SendHisMsg(session *sessions.FrontSession) {
	romMutex.Lock()
	defer romMutex.Unlock()

	chatroom := findChatroom(session.RId)
	if chatroom == nil {
		return
	}

	fmt.Println(chatroom.msgMap)
	len := len(chatroom.msgMap)
	speakerId := make([]uint64, len)
	speakerName := make([]string, len)
	speakerText := make([]string, len)
	time := make([]string, len)

	for i, msg := range chatroom.msgMap {
		speakerId[i] = msg.Rid
		speakerId[i] = msg.Rid
		speakerName[i] = msg.Rname
		speakerText[i] = msg.Text
		time[i] = msg.Time
	}

	msgBody := &gameProto.ClientChatHistoryS2C{
		SpeakerId:   speakerId,
		SpeakerName: speakerName,
		SpeakerText: speakerText,
		Time:        time,
	}
	send.SendMsgToClient(session, msgBody)
}

func Broadcast(rid uint64, msgBody proto.Message) {
	romMutex.Lock()
	defer romMutex.Unlock()

	chatroom := findChatroom(rid)
	if chatroom != nil {
		for rid, _ := range chatroom.joinPlayerMap {
			session := sessions.GetFrontSession(rid)
			send.SendMsgToClient(session, msgBody)
		}
	}
}

func NewChatRoom() *ChatRooms {
	room := &ChatRooms{
		roomId:        guid_.NewID(),
		roomName:      "chatroom",
		joinPlayerMap: make(map[uint64]joinPlayers),
		time :         time.Now().Unix(),
		msgMap    :    make([]Msg, 0, maxMsgMap+1),
		popular: NewPopular(),
	}
	roomMap[room.roomId] = room
	return room
}

func getDefaultRoom() *ChatRooms {
	for _, chatroom := range roomMap{
		return chatroom
	}
	return NewChatRoom()
}

func getChatroom(rid uint64) *ChatRooms {
	if chatroom, ok := ridJoinMap[rid]; ok {
		return chatroom
	}else {
		return getDefaultRoom()
	}
}

func JoinRoom(rid uint64) {
	romMutex.Lock()
	defer romMutex.Unlock()

	session := sessions.GetFrontSession(rid)
	if session != nil {
		chatroom := getChatroom(rid)
		joinPlayer := joinPlayers{
			joinTime: time.Now(),
			rid:      session.RId,
			rname:    session.RName,
		}
		chatroom.joinPlayerMap[rid] = joinPlayer
		ridJoinMap[rid]  = chatroom
		rnameJoinMap[session.RName]=rid
		session.SetChat(room)
		log.Println("JoinRoom", session.RId, session.RName, chatroom.roomId)
	}
}

func (room rooms) QuitRoom(rid uint64) {
	romMutex.Lock()
	defer romMutex.Unlock()

	if chatroom, ok := ridJoinMap[rid]; ok {
		delete(rnameJoinMap, chatroom.joinPlayerMap[rid].rname )
		delete(chatroom.joinPlayerMap, rid)
	}
	delete(ridJoinMap, rid)
	log.Println("QuitRoom", rid)
}

func findChatroom(rid uint64) *ChatRooms {
	if chatroom, ok_ := ridJoinMap[rid]; ok_ {
		return chatroom
	}
	return nil
}

func RegMsg(session *sessions.FrontSession, text string) {
	romMutex.Lock()
	defer romMutex.Unlock()

	chatroom := findChatroom(session.RId)
	if chatroom != nil {
		chatroom.msgMap = append(chatroom.msgMap, packMsg(session, text))
		if len(chatroom.msgMap) > maxMsgMap {
			chatroom.msgMap = chatroom.msgMap[1:]
		}
		fmt.Println(chatroom.msgMap)
		chatroom.popular.AddPopular(text)
	}
}


func GetPopular(rid uint64, n uint32) string {
	chatroom := findChatroom(rid)
	if chatroom != nil {
		return chatroom.popular.GetPopular(int64(n))
	}else {
		return ""
	}
}