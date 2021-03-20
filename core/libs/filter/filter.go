package filter

import (
	"GoGame/core/libs/trie"
	"bufio"
	"io"
	"os"
)

type Filter struct {
	trie  *trie.Trie
}
func New() *Filter {
	return &Filter{
		trie: trie.NewTrie(),
	}
}

func (filter *Filter) LoadWordDict(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return filter.Load(f)
}

func (filter *Filter) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		filter.trie.Add(string(line))
	}

	return nil
}

func (this *Filter) Replace(msg string, character rune) string {
	return this.trie.Replace(msg,character)
}