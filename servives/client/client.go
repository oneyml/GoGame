package main

import (
	"GoGame/core/config"
	"GoGame/core/protos"
	"GoGame/core/protos/gameProto"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

var (
	client       = new(clientSession)
)

type clientSession struct {
	con       *websocket.Conn
	account   string
	token     string
	receMutx  sync.Mutex
	sendMutx  sync.Mutex
	closeFlag int32
	rid       uint64
	rname     string
}


func ClientStart() {
	go startConnect()
}

func startConnect() {
	u := url.URL{Scheme: "ws", Host: config.GetServerIp(), Path: "/"}
	d := websocket.DefaultDialer
	c, _, err := d.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("startConnect:", err)
		return
	}
	client.con   = c
	client.rid   = Args.Rid
	client.rname = Args.Rname
	go client.recvMsg()
	client.login()
	client.ping()
}

//是否关闭
func (this *clientSession) isClose() bool {
	return atomic.LoadInt32(&this.closeFlag) == 1
}

func (this *clientSession) close() {
	if atomic.CompareAndSwapInt32(&this.closeFlag, 0, 1) {
		this.con.Close()
	}
}

func (this *clientSession) recvMsg() {
	for true {
		if this.isClose() {
			break
		}
		_, message, err := this.con.ReadMessage()
		if err != nil {
			fmt.Println("recvMsg:", err)
			break
		}
		msgBody := message[2:]
		protoMsg := protos.UnMarshalProtoMsg(msgBody)
		this.handleMsg(protoMsg)
	}
	this.close()
}

//改成注册的方式
func (this *clientSession) handleMsg(protoMsg protos.ProtoMsg) {
	if protoMsg.MsgId == gameProto.ID_client_login_s2c {
		loginS2C(protoMsg.MsgBody)
	} else if protoMsg.MsgId == gameProto.ID_client_chatJoin_s2c {
	} else if protoMsg.MsgId == gameProto.ID_client_chatHistory_s2c {
		chatShowHis(protoMsg.MsgBody)
	} else if protoMsg.MsgId == gameProto.ID_client_chat_s2c {
		chatShowChat(protoMsg.MsgBody)
	} else if protoMsg.MsgId == gameProto.ID_client_chatStats_s2c {
		chatShowStats(protoMsg.MsgBody)
	} else if protoMsg.MsgId == gameProto.ID_client_chatPopular_s2c {
		chatShowPopular(protoMsg.MsgBody)
	}
}

func (this *clientSession) sendMsg(args proto.Message) {
	if this.isClose() {
		return
	}

	args1 := protos.MarshalProtoMsg(args)
	msgLen := uint16(2 + len(args1))
	result := make([]byte, msgLen)
	binary.BigEndian.PutUint16(result[:2], msgLen)
	copy(result[2:], args1)
	this.con.WriteMessage(websocket.BinaryMessage, result)
}

func GetClient() *clientSession {
	return client
}

func (this *clientSession) login() {
	msgBody := &gameProto.ClientLoginC2S{
		RId:   proto.Uint64(this.rid),
		RName: proto.String(this.rname),
	}
	this.sendMsg(msgBody)
}

func loginS2C(message proto.Message)  {
	data := message.(*gameProto.ClientLoginS2C)
	if data.GetError() == 1 {
		client := GetClient()
		log.Println("exist player name:", client.rname)
	}
}

func (this *clientSession) ping() {
	for true {
		time.Sleep(time.Second * 5)
		msg := &gameProto.ClientPingC2S{}
		this.sendMsg(msg)
	}
}