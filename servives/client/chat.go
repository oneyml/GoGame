package main

import (
	"GoGame/core/config"
	"GoGame/core/protos/gameProto"
	"github.com/golang/protobuf/proto"
	"log"
	"strconv"
)

func chatJoin() {
	client := GetClient()
	msgBody := &gameProto.ClientChatJoinC2S{
		RId: &client.rid,
	}
	client.sendMsg(msgBody)
}

func chatChat(text string) {
	client := GetClient()
	msgBody := &gameProto.ClientChatC2S{
		Msg: proto.String(text),
	}
	client.sendMsg(msgBody)
}

func chatStats(name string) {
	client := GetClient()
	msgBody := &gameProto.ClientChatStatsC2S{
		Rname: proto.String(name),
	}
	client.sendMsg(msgBody)
}
func chatPopular(n string) {
	n_, _:= strconv.Atoi(n)
	if n_ >= 60 {
		log.Println(" n need lees than:%d", config.GetChatPopular())
	} else {
		client := GetClient()
		msgBody := &gameProto.ClientChatPopularC2S{
			N: proto.Uint32(uint32(n_)),
		}
		client.sendMsg(msgBody)
	}
}

func chatShowHis(message proto.Message) {
	log.Println("HISTORY INFO")
	data := message.(*gameProto.ClientChatHistoryS2C)
	for i, _ := range data.SpeakerId {
		log.Printf("i:%d, id:%d, name:%s, text:%s, time:%s\n",
			i, data.SpeakerId[i], data.SpeakerName[i], data.SpeakerText[i], data.Time[i])
	}
}

func chatShowChat(message proto.Message) {
	data := message.(*gameProto.ClientChatS2C)
	log.Printf("id:%d, name:%s, text:%s\n",
		data.GetSpeakerId(), data.GetSpeakerName(), data.GetSpeakerText())
}
func chatShowStats(message proto.Message) {
	data := message.(*gameProto.ClientChatStatsS2C)
	if len(data.GetTime()) > 0 {
		log.Printf("name:%s\nname:%s\n",
			data.GetRname(), data.GetTime())
	} else {
		log.Printf("not find player:%s", data.GetRname())
	}
}

func chatShowPopular(message proto.Message) {
	data := message.(*gameProto.ClientChatPopularS2C)
	log.Printf("%s", data.GetWorld())
}