package service

import "GoGame/core/libs/websocket"

type Service struct {
	name     string
	id       int
	ports    map[string]string
	wsServer *websocket.Server
}

func NewService(name string) *Service {
	service := &Service{
		name:     name,
		id:       0,
		ports:    nil,
		wsServer: nil,
	}
	return service
}

func (this *Service) StartWebSocket(port string) {
	this.wsServer = websocket.NewServer(port)
	this.wsServer.Start()
}
