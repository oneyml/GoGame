package sessions

import (
	"encoding/binary"
	"github.com/gorilla/websocket"
	"log"
)

func NewFrontCodec(rw *websocket.Conn) *frontByteCodec {
	return &frontByteCodec{
		rw: rw,
	}
}

type frontByteCodec struct {
	rw *websocket.Conn
}

func (this *frontByteCodec) Receive() ([]byte, error) {
	_, msg, err := this.rw.ReadMessage()
	if err != nil {
		return nil, err
	}
	if len(msg) < 2 {
		log.Println("front receive message len error")
		return nil, nil
	}
	return msg[2:], nil
}

func (this *frontByteCodec) Send(args1 []byte) error {
	msgLen := uint16(2 + len(args1))
	result := make([]byte, msgLen)
	binary.BigEndian.PutUint16(result[:2], msgLen)
	copy(result[2:], args1)
	this.rw.WriteMessage(websocket.BinaryMessage, result)
	return nil
}

func (this *frontByteCodec) Close() error {
	this.rw.Close()
	return nil
}
