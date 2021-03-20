package chatroom

import (
	"fmt"
	"testing"
	"time"
)

func TestPopular(t *testing.T) {
	popular := NewPopular()
	popular.AddPopular("hello world")
	time.Sleep(time.Second)
	popular.AddPopular("hello world")
	time.Sleep(time.Second)
	popular.AddPopular("world")
	time.Sleep(3*time.Second)
	fmt.Println(popular.GetPopular(10))
}