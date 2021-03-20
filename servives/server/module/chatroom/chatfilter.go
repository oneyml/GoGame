package chatroom

import "GoGame/core/libs/filter"

var (
	chatFilter = filter.New()
)

func init()  {
	chatFilter.LoadWordDict("./../../../../cfg/list.txt")
}

func Filter(msg string) string {
	return chatFilter.Replace(msg, '*')
}