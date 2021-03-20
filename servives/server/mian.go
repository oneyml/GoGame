package main

import (
	"GoGame/core/config"
	"GoGame/core/messages"
	"GoGame/core/protos/gameProto"
	"GoGame/core/service"
	"GoGame/servives/server/module"
	"time"
)

func initMsgModule() {
	messages.RegisterIpcServerHandle(gameProto.ID_client_ping_c2s, module.Ping)
	messages.RegisterIpcServerHandle(gameProto.ID_client_login_c2s, module.Login)

	messages.RegisterIpcServerHandle(gameProto.ID_client_chat_c2s,     module.Chat)
	messages.RegisterIpcServerHandle(gameProto.ID_client_chatJoin_c2s, module.ChatJoin)

	messages.RegisterIpcServerHandle(gameProto.ID_client_chatPopular_c2s, module.Popular)
	messages.RegisterIpcServerHandle(gameProto.ID_client_chatStats_c2s,   module.Stats)
}
func main() {
	service := service.NewService("game")
	service.StartWebSocket(config.GetListenPort())
	initMsgModule()
	for true {
		time.Sleep(time.Second*10)
	}
}
