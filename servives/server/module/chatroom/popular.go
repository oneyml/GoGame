package chatroom

import (
	"GoGame/core/config"
	"fmt"
	"strings"
	"time"
)

var (
	expire = int64(config.GetChatPopular())
)

type populars struct {
	//countMap map[string]int
	timeMap  map[int64]map[string]int
}

func NewPopular() *populars {
	return &populars{
		//countMap:    make(map[string]int),
		timeMap: 	 map[int64]map[string]int{},
	}
}

/*
func (this*populars)addCount(node []string)  {
	for _, word := range node {
		if _, ok := this.countMap[word]; ok {
			this.countMap[word] = this.countMap[word] + 1
		}else{
			this.countMap[word] = 1
		}
	}
}

func (this*populars)delCount(node map[string]int)  {
	for word, count := range node {
		if _, ok := this.countMap[word]; ok {
			oldCount := this.countMap[word]
			if oldCount != count {
				this.countMap[word] = this.countMap[word]	- count
			}else{
				delete(this.countMap, word)
			}
		}
	}
}
*/

func (this*populars)deleteExpireNode()  {
	now := time.Now().Unix()
	for last, _:= range this.timeMap {
		if now - last > expire {
			delete(this.timeMap, last)
			//this.delCount(node)
		}
	}
}

func (this*populars)addNewNode(textArray []string)  {
	now := time.Now().Unix()
	if this.timeMap[now] == nil {
		this.timeMap[now] = make(map[string]int)
	}
	node := this.timeMap[now]
	for _, word := range textArray {
		if _, ok := node[word]; ok {
			node[word] = node[word] + 1
		}else{
			node[word] = 1
		}
	}
	fmt.Println(this.timeMap)
}

func (this*populars)popular(n int64) string {
	i := 0
	w := ""
	countNode := make(map[string]int)
	now := time.Now().Unix()
	for last, node := range this.timeMap {
		if now - last <= n {
			for word, cnt := range node {
				if _, ok := countNode[word]; ok {
					countNode[word] += cnt
				} else {
					countNode[word] = 1
				}
				if i < countNode[word] {
					i = countNode[word]
					w = word
				}
			}
		}
	}
	return w
}

func (this*populars)AddPopular(text string)  {
	textArray := strings.Split(text, " ")
	this.deleteExpireNode()
	this.addNewNode(textArray)
	//this.addCount(textArray)
}

func (this*populars)GetPopular(n int64) string {
	this.deleteExpireNode()
	return this.popular(n)
}