package implement_trie_prefix_tree

type Trie struct {
    children [26]*Trie
    isWord bool
}

func Constructor() Trie {
    return Trie{[26]*Trie{}, false}
}

func (this *Trie) get(letter byte) *Trie {
    return this.children[letter - 'a']
}

func (this *Trie) Insert(word string) {
    for i := 0; i < len(word); this, i = this.get(word[i]), i + 1 {
        if this.get(word[i]) == nil {
            this.children[word[i] - 'a'] = &Trie{[26]*Trie{}, false}
        }
    }

    this.isWord = true
}

func (this *Trie) Search(word string) bool {
    for i := 0; i < len(word) && this != nil; this, i = this.get(word[i]), i + 1 { }
    return this != nil && this.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
    for i := 0; i < len(prefix) && this != nil; this, i = this.get(prefix[i]), i + 1 { }
    return this != nil
}