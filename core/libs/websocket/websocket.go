package websocket

import (
	"GoGame/core/libs/sessions"
	"GoGame/core/messages"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SessionCreateHandle func(session *sessions.FrontSession)
type SessionMsgHandle func(session *sessions.FrontSession, msgBody []byte)

type Server struct {
	port   string

	sessionCreatHandle SessionCreateHandle
	sessionMsgHandle   SessionMsgHandle
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (this *Server) Start() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				fmt.Println("http.HandleFunc")
				return
			}
			codec := sessions.NewFrontCodec(conn)
			clientSession := sessions.NewSession(conn, codec)
			clientSession.SetMsgHandle(messages.MsgHandle)

			// sessions.AddFrontSession(clientSession)
			// 处理消息
			this.recvie(clientSession)
			// 收消息
			//clientSession.loop()
		})

		err := http.ListenAndServe(this.port, nil)
		if err != nil {
			fmt.Println("http.ListenAndServer", err)
		}
	}()
}
func (this *Server) recvie(session *sessions.FrontSession) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
		fmt.Println("recive occer close")
		session.Close()
	}()
	for true {
		msg, err := session.Receive()
		if err != nil || msg == nil {
			fmt.Println("reciv eror", msg, err)
			break
		}
	}
}
