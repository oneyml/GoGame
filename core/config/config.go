package config

//服务监听端口
func GetListenPort() string {
	return ":7777"
}
//客户端连接ip
func GetServerIp() string {
	return "127.0.0.1:7777"
}

//最受欢迎的单词过期时间
func GetChatPopular() uint8 {
	return 60
}

//聊天历史信息条数
func GetChatHistory() uint8 {
	return 60
}
