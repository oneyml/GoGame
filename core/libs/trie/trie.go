package trie

// Trie树节点.
type Node struct {
	isRootNode bool
	isWordEnd  bool
	character  rune
	children   map[rune]*Node
}

// Trie树.
type Trie struct {
	Root *Node
}

// 根节点
func NewRootNode(character rune) *Node {
	return &Node{
		isRootNode: true,
		character:  character,
		children:   make(map[rune]*Node, 0),
	}
}

// 新建Trie树
func NewTrie() *Trie {
	return &Trie{
		Root: NewRootNode(0),
	}
}

// 子节点
func NewNode(character rune) *Node {
	return &Node{
		character: character,
		children:  make(map[rune]*Node, 0),
	}
}

// 添加一个单词
func (tree *Trie) Add(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for i:= 0; i< len(runes); i++ {
		r := runes[i]
		if next, ok := current.children[r]; ok {
			current = next
		} else {
			newNode := NewNode(r)
			current.children[r] = newNode
			current = newNode
		}
		if i== len(runes)-1 {
			current.isWordEnd = true
		}
	}
}

// 词语替换 *
func (tree *Trie) Replace(text string, character rune) string {
	var (
		parent  = tree.Root
		current *Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < len(runes); position++ {
		current, found = parent.children[runes[position]]

		if !found || (!current.isWord() && position == length-1) {
			parent = tree.Root
			position = left
			left++
			continue
		}

		if current.isWord() && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = character
			}
		}

		parent = current
	}

	return string(runes)
}

// 判断是否为某个单词的结束
func (node *Node) isWord() bool {
	return node.isWordEnd
}