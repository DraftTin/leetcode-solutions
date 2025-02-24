package main

type TrieNode211 struct {
	Children [26]*TrieNode211
	isWord   bool
}

type WordDictionary struct {
	root *TrieNode211
}

func Constructor() WordDictionary {
	dict := WordDictionary{}
	dict.root = &TrieNode211{}
	return dict
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, c := range word {
		if cur.Children[c - 'a'] == nil {
			cur.Children[c - 'a'] = &TrieNode211{}
		}
		cur = 
	}
}

func (this *WordDictionary) Search(word string) bool {

}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
