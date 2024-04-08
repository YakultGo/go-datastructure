package datastructure

type Trie struct {
    nxt [26]*Trie
    tag int
}

func (this *Trie) Insert(word string) {
    root := this
    for _, c := range word {
        if root.nxt[c-'a'] == nil {
            root.nxt[c-'a'] = &Trie{}
        }
        root = root.nxt[c-'a']
    }
    root.tag = 1
}

func (this *Trie) Search(word string) bool {
    root := this
    for _, c := range word {
        if root.nxt[c-'a'] == nil {
            return false
        }
        root = root.nxt[c-'a']
    }
    return root.tag == 1
}

func (this *Trie) StartsWith(prefix string) bool {
    root := this
    for _, c := range prefix {
        if root.nxt[c-'a'] == nil {
            return false
        }
        root = root.nxt[c-'a']
    }
    return true
}
