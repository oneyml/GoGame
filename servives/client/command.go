package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// 处理命令
// /chat
// /popular
// //stats
const (
	command_join    = "/join"
	command_chat    = "/chat"
	command_popular = "/popular"
	command_stats   = "/stats"
)

func CommandStart() {
	go func() {
		for true {
			reader := bufio.NewReader(os.Stdin)
			cmdString, err := reader.ReadString('\n')
			if err != nil {
				log.Println("error", err)
				continue
			} else {
				cmdString = strings.Trim(cmdString, "\n")
			}
			commands := strings.SplitN(cmdString, " ", 2)
			command := commands[0]
			if command == command_join {
				chatJoin()
			} else if command == command_chat {
				if len(commands) > 1 {
					chatChat(commands[1])
				}
			} else if command == command_popular {
				if len(commands) > 1 {
					chatPopular(commands[1])
				} else {
					log.Println("need n")
				}
			} else if command == command_stats {
				if len(commands) > 1 {
					chatStats(commands[1])
				}
			}
		}
	}()
}
