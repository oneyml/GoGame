package gameProto

const (
	ID_client_ping_c2s  = 1001
	ID_client_login_c2s = 1002
	ID_client_login_s2c = 1003

	//加入聊天C2S(2001)
	ID_client_chatJoin_c2s = 2001
	//加入聊天S2C(2002)
	ID_client_chatJoin_s2c = 2002
	//聊天消息C2S(2003)
	ID_client_chat_c2s = 2003
	//聊天消息S2C(2004)
	ID_client_chat_s2c = 2004
	//聊天历史S2C(2005)
	ID_client_chatHistory_s2c = 2005

	//聊天popular(2006)
	ID_client_chatPopular_c2s = 2006

	//聊天popular(2007)
	ID_client_chatPopular_s2c = 2007

	//聊天stats(2008)
	ID_client_chatStats_c2s = 2008

	//聊天popular(2009)
	ID_client_chatStats_s2c = 2009
)
