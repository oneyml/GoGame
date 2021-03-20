package sessions

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type Codec interface {
	Receive() ([]byte, error)
	Send([]byte) error
	Close() error
}

type Chat interface {
	QuitRoom(uint642 uint64)
}

type handleMsg func(session *FrontSession, msgBody []byte)
type FrontSession struct {
	conn      *websocket.Conn
	recvChan  chan []byte
	msgHandle handleMsg
	codec     Codec
	recvMutex sync.Mutex
	sendMutex sync.Mutex
	closeFlag int32
	RId       uint64
	RName     string
	chat      Chat
	Time      int64
}

func NewSession(conn *websocket.Conn, codec Codec) *FrontSession {
	clientSession := FrontSession{
		conn      : conn,
		codec     : codec,
		recvChan  : make(chan []byte, 100),
		Time 	  : time.Now().Unix(),
	}
	clientSession.loop()
	return &clientSession
}

func (this *FrontSession) SetChat(chat Chat) {
	this.chat = chat
}

func (this *FrontSession) SetMsgHandle(msgHandle handleMsg) {
	this.msgHandle = msgHandle
}

func (this *FrontSession) IsClosed() bool {
	return atomic.LoadInt32(&this.closeFlag) == 1
}

func (this *FrontSession) Close() {
	if atomic.CompareAndSwapInt32(&this.closeFlag, 0, 1) {
		this.recvMutex.Lock()
		log.Println("session close")
		close(this.recvChan)
		for _ = range this.recvChan {
		}
		this.recvMutex.Unlock()
		if this.chat != nil {
			this.chat.QuitRoom(this.RId)
		}
		RemoveFrontSession(this.RId)
		this.codec.Close()
	}
}

func (this *FrontSession) Receive() ([]byte, error) {
	msgBody, err := this.codec.Receive()
	if msgBody != nil {
		this.recvMutex.Lock()
		if this.IsClosed() {
			this.recvMutex.Unlock()
			return nil, err
		}
		this.recvChan <- msgBody
		this.recvMutex.Unlock()
	}
	return msgBody, err
}
var ErrClosed = errors.New("session closed")
func (this *FrontSession) Send(msg []byte) (err error) {
	if this.IsClosed() {
		return ErrClosed
	}

	this.sendMutex.Lock()
	defer this.sendMutex.Unlock()

	return this.codec.Send(msg)
}

func (this *FrontSession) loop() {
	defer func() {
		if x := recover(); x != nil {
			log.Println("loop recover", x)
		}
	}()
	go func() {
		for true {
			select {
			// 这个地方的语法理解有点问题
			case data, err := <-this.recvChan:
				if data != nil {
					this.msgHandle(this, data)
				} else {
					log.Println("frontSession loop recvChan error", err)
					return
				}
			}
		}
	}()
}
