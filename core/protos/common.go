package protos

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

var (
	msgObjectMap = make(map[uint16]reflect.Type)
	msgIDMap     = make(map[reflect.Type]uint16)
)

func SetMsg(msgId uint16, data interface{}) {
	msgType := reflect.TypeOf(data)
	msgObjectMap[msgId] = msgType
	msgIDMap[reflect.TypeOf(reflect.New(msgType).Interface())] = msgId
}

// 通过消息id获取该消息的结构体
func GetMsgObject(msgID uint16) proto.Message {
	if msgType, exists := msgObjectMap[msgID]; exists {
		return reflect.New(msgType).Interface().(proto.Message)
	} else {
		fmt.Println("No register msgObjectMap", msgID)
		return nil
	}
}

func GetMsgId(msg interface{}) uint16 {
	msgType := reflect.TypeOf(msg)
	if msgId, exists := msgIDMap[msgType]; exists {
		return msgId
	}
	return 0
}

func MarshalProtoMsg(msg proto.Message) []byte {
	msg1, _ := proto.Marshal(msg)
	result := make([]byte, 2+len(msg1))
	msgId := GetMsgId(msg)
	binary.BigEndian.PutUint16(result[:2], msgId)
	copy(result[2:], msg1)
	return result
}

func UnMarshalProtoMsgId(msg []byte) uint16 {
	return binary.BigEndian.Uint16(msg[:2])
}

type ProtoMsg struct {
	MsgId   uint16
	MsgBody proto.Message
}

func UnMarshalProtoMsg(msg []byte) ProtoMsg {
	msgId := UnMarshalProtoMsgId(msg)
	msgBody := GetMsgObject(msgId)
	proto.Unmarshal(msg[2:], msgBody)
	return ProtoMsg{
		MsgId:   msgId,
		MsgBody: msgBody,
	}
}
