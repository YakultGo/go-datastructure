type Trie struct {
	nxt [26]*Trie
	tag int
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

func (tr *Trie) NewTrie() *Trie {
	return &Trie{}
}