package main

type TrieNode208 struct {
	Children [26]*TrieNode208
	IsWord   bool
}

type Trie struct {
	Root *TrieNode208
}

func Constructor() Trie {
	trie := Trie{}
	trie.Root = &TrieNode208{}
	return trie
}

func (this *Trie) Insert(word string) {
	cur := this.Root
	for _, c := range word {
		if cur.Children[c-'a'] == nil {
			cur.Children[c-'a'] = &TrieNode208{}
		}
		cur = cur.Children[c-'a']
	}
	cur.IsWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this.Root
	for _, c := range word {
		if cur.Children[c-'a'] == nil {
			return false
		}
		cur = cur.Children[c-'a']
	}
	return cur.IsWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.Root
	for _, c := range prefix {
		if cur.Children[c-'a'] == nil {
			return false
		}
		cur = cur.Children[c-'a']
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
