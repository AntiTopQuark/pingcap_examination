package Process

import (
	"math"
)

type Node struct {
	Count      uint64
	FirstIndex uint64
}
type Collection interface {
	Add(word string, index uint64)
	GetResult() (string, uint64)
}

type MapCollection struct {
	data map[string]*Node
}

func (receiver *MapCollection) Add(word string, index uint64) {
	value, ok := receiver.data[word]
	if ok {
		// 该单词存在
		value.Count += 1
		if index < value.FirstIndex {
			value.FirstIndex = index
		}
	} else {
		receiver.data[word] = &Node{Count: 1, FirstIndex: index}
	}
}

func (receiver *MapCollection) GetResult() (string, uint64) {
	var word string
	var firstIndex uint64 = math.MaxUint64
	for key := range receiver.data {
		node := receiver.data[key]
		if node.Count == 1 && node.FirstIndex < firstIndex {
			word = key
			firstIndex = node.FirstIndex
		}
	}
	return word, firstIndex
}

// 字典树
type Trie struct {
	isEnd    bool
	children [26]*Trie
	value    *Node
}

func Constructor() Trie {
	return Trie{}
}

func (receiver *Trie) Insert(word string, index uint64) {
	cur := receiver
	for i, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			cur.children[n] = &Trie{}

		}
		cur = cur.children[n]
		if i == len(word)-1 {
			cur.isEnd = true
			if cur.value == nil {
				cur.value = &Node{Count: 1, FirstIndex: index}
			} else {
				cur.value.Count++
				if index < cur.value.FirstIndex {
					cur.value.FirstIndex = index
				}
			}

		}
	}
}
func (receiver *Trie) Search(key string, Result_Word *string, Result_FirstIndex *uint64) {
	cur := receiver

	if cur.isEnd && cur.value.Count == 1 {

		if cur.value.FirstIndex < *Result_FirstIndex {
			*Result_Word = key
			*Result_FirstIndex = cur.value.FirstIndex
		}
	}
	for i := 0; i < len(cur.children); i++ {
		if cur.children[i] == nil {
			continue
		} else {
			var _char byte
			_char = byte('a' + i)

			cur.children[i].Search(key+string(_char), Result_Word, Result_FirstIndex)
		}
	}
}

type TrieCollection struct {
	data Trie
}

func (receiver *TrieCollection) Add(word string, index uint64) {
	receiver.data.Insert(word, index)
}

func (receiver *TrieCollection) GetResult() (string, uint64) {
	var Result_Word string
	var Result_firstIndex uint64 = math.MaxUint64
	receiver.data.Search("", &Result_Word, &Result_firstIndex)
	return Result_Word, Result_firstIndex
}
