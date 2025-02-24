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
		if cur.Children[c-'a'] == nil {
			cur.Children[c-'a'] = &TrieNode211{}
		}
		cur = cur.Children[c-'a']
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	cur := this.root
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			if cur.Children[word[i]-'a'] == nil {
				return false
			}
			cur = cur.Children[word[i]-'a']
		} else {
			for j := 0; j < 26; j++ {
				if cur.Children[j] != nil {
					newDict := &WordDictionary{root: cur.Children[j]}
					if newDict.Search(word[i+1:]) {
						return true
					}
				}
			}
			return false
		}
	}
	return cur.isWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
