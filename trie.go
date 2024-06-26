package datastructure

import "runtime/debug"

// 禁止GC，能够提高性能（参考灵神）
func init() { debug.SetGCPercent(-1) }

type Trie struct {
    nxt [26]*Trie
    tag int
}

func (tr *Trie) NewTrie() *Trie {
    return &Trie{}
}
func (tr *Trie) Insert(word string) {
    root := tr
    for _, c := range word {
        if root.nxt[c-'a'] == nil {
            root.nxt[c-'a'] = &Trie{}
        }
        root = root.nxt[c-'a']
    }
    root.tag = 1
}

func (tr *Trie) Search(word string) bool {
    root := tr
    for _, c := range word {
        if root.nxt[c-'a'] == nil {
            return false
        }
        root = root.nxt[c-'a']
    }
    return root.tag == 1
}

func (tr *Trie) StartsWith(prefix string) bool {
    root := tr
    for _, c := range prefix {
        if root.nxt[c-'a'] == nil {
            return false
        }
        root = root.nxt[c-'a']
    }
    return true
}
