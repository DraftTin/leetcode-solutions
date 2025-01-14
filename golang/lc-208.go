type TrieNode struct {
	Children [26]*TrieNode
	Val      string
	IsPrefix bool
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	trie := Trie{}
	trie.Root = &TrieNode{}
	return trie
}

func (this *Trie) Insert(word string) {
	node := this.Root
	for _, c := range word {
		if node.Children[c-'a'] == nil {
			node.Children[c-'a'] = &TrieNode{Val: node.Val + string(c), IsPrefix: true}
		}
		node = node.Children[c-'a']
	}
	node.IsPrefix = false
}

func (this *Trie) Search(word string) bool {
	node := this.Root
	for _, c := range word {
		if node.Children[c-'a'] == nil {
			return false
		}
		node = node.Children[c-'a']
	}
	return !node.IsPrefix
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.Root
	for _, c := range prefix {
		if node.Children[c-'a'] == nil {
			return false
		}
		node = node.Children[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
