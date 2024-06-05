package main

import (
	"sync"
)

// Trie 节点结构
type TrieNode struct {
	children  map[rune]*TrieNode
	isEnd     bool
	frequency int
}

// Trie 树结构
type Trie struct {
	root *TrieNode
	mu   sync.Mutex
}

// 创建新的 Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// 向 Trie 中插入单词
func (t *Trie) Insert(word string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.frequency++
}

// 查找单词在 Trie 中的频率
func (t *Trie) Search(word string) int {
	t.mu.Lock()
	defer t.mu.Unlock()

	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return 0
		}
		node = node.children[ch]
	}
	if node.isEnd {
		return node.frequency
	}
	return 0
}

//该方法仅用作二叉树方法测试，
//func main() {
//	trie := NewTrie()
//	words := []string{"hello", "world", "hello", "trie"}
//	for _, word := range words {
//		trie.Insert(word)
//	}
//
//	fmt.Printf("Frequency of 'hello': %d\n", trie.Search("hello"))
//	fmt.Printf("Frequency of 'world': %d\n", trie.Search("world"))
//	fmt.Printf("Frequency of 'trie': %d\n", trie.Search("trie"))
//}
