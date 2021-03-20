package filter

import (
	"GoGame/core/libs/trie"
	"fmt"
	"os"
	"testing"
)

func TestTrieTree(t *testing.T) {
	tree := trie.NewTrie()
	tree.Add("中")
	tree.Add("西塞罗")
	fmt.Println(tree.Replace("阅读中的少年西塞罗", '*'))
}

func TestLoadDict(t *testing.T) {
	filter := New()
	dir, _ := os.Getwd()
	fmt.Println(dir)
	err := filter.LoadWordDict("./../../../cfg/list.txt")
	if err != nil {
		t.Errorf("fail to load dict %v", err)
	}

	fmt.Println(filter.trie.Replace("阅读中的少年xxx西塞罗", '*'))
}



