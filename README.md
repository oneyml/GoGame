#客服端代码文件
```
client
    mian.go     //启动文件
    command.go  //获取命令行输入
    client.go   //连接 发生 接收
```

#服务器代码文件
```
chatroom          //聊天室
    chatfilter.go //单词过滤
    popular.go    //最流行单词
    room.go       //聊天房间 历史信息
chat.go           //聊天
login.go          //登录
ping.go           //
stats.go          //玩家在线时长
mian.go           //启动文件
```

#前后端通信
    协议使用 google proto
    协议文件 gameProto.proto

# 过滤词算法
    trie树 trie.go
# popular算法
    timeMap  map[int64]map[string]int
    把每秒的单词进行统计并放入timeMap数据结构
    timeMap  的key是unix时间戳
    
# 配置文件 config.go
    服务监听端口 
    客户端连接ip
    
    最受欢迎的单词过期时间
    聊天历史信息条数
    

# 启动 在bin目录下
```
   // 服务器
   server.exe 
    
   // 客户端
   client.exe -i 1001 -n name1
   -i 是玩家id,如果重复会把对方踢下线
   -n 是玩家名字,如果名字不唯一需重新登录
```

# 使用说明
```
加入聊天室，并获得历史信息
    /join
    
聊天命令 需要先加入聊天室才可以聊天
    /chat hello world
最受欢迎单词
    /popular 10
玩家时常
    /stats name1
```
