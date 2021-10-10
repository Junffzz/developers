/*
@Time : 2021/10/6 11:09
@Author : ZhaoJunfeng
@File : trie_test
*/
package basic

import "testing"

func TestTrie(t *testing.T) {
    word:="abcdef"
    obj := TrieConstructor()
    obj.Insert(word)
    param_2 := obj.Search(word)
    // param_3 := obj.StartsWith(prefix)
    t.Logf("%v\n",param_2)
}

type Trie struct {
    children [26]*Trie
    isEnd    bool
}

func TrieConstructor() Trie {
    return Trie{}
}

func (t *Trie) Insert(word string) {
    node := t
    for _, ch := range word {
        ch -= 'a'
        if node.children[ch] == nil {
            node.children[ch] = &Trie{}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
    node := t
    for _, ch := range prefix {
        ch -= 'a'
        if node.children[ch] == nil {
            return nil
        }
        node = node.children[ch]
    }
    return node
}

func (t *Trie) Search(word string) bool {
    node := t.SearchPrefix(word)
    return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    return t.SearchPrefix(prefix) != nil
}

