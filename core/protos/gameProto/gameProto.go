package gameProto

import "GoGame/core/protos"

func init() {
	protos.SetMsg(ID_client_ping_c2s, ClientPingC2S{})
	protos.SetMsg(ID_client_login_c2s, ClientLoginC2S{})
	protos.SetMsg(ID_client_login_s2c, ClientLoginS2C{})

	protos.SetMsg(ID_client_chatJoin_c2s, ClientChatJoinC2S{})
	protos.SetMsg(ID_client_chatJoin_s2c, ClientChatJoinS2C{})

	protos.SetMsg(ID_client_chat_c2s, ClientChatC2S{})
	protos.SetMsg(ID_client_chat_s2c, ClientChatS2C{})

	protos.SetMsg(ID_client_chatHistory_s2c, ClientChatHistoryS2C{})

	protos.SetMsg(ID_client_chatPopular_c2s, ClientChatPopularC2S{})
	protos.SetMsg(ID_client_chatPopular_s2c, ClientChatPopularS2C{})

	protos.SetMsg(ID_client_chatStats_c2s, ClientChatStatsC2S{})
	protos.SetMsg(ID_client_chatStats_s2c, ClientChatStatsS2C{})

}
