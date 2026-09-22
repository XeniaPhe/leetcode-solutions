package design_add_and_search_words_data_structure

type WordDictionary struct {
    children [26]*WordDictionary
    isWord bool
}

func Constructor() WordDictionary {
    return WordDictionary{[26]*WordDictionary{}, false}
}

func (this *WordDictionary) AddWord(word string)  {
    for i := 0; i < len(word); this, i = this.get(word[i]), i + 1 {
        if this.get(word[i]) == nil {
            this.children[word[i] - 'a'] = &WordDictionary{[26]*WordDictionary{}, false}
        }
    }

    this.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
    type stackType struct {
        trie *WordDictionary
        wordIdx, childIdx int
    }

    for stack, curr, currIdx := []stackType{}, this, 0; ; {
        for ; curr != nil && currIdx < len(word); currIdx += 1 {
            if letter := word[currIdx]; letter == '.' {
                stack, curr = append(stack, stackType{curr, currIdx + 1, 0}), nil
            } else {
                curr = curr.get(letter)
            }
        }

        if curr != nil && curr.isWord {
            return true
        } else if len(stack) == 0 {
            return false
        } else if idx, top := len(stack) - 1, stack[len(stack) - 1]; true {
            curr, currIdx = top.trie.children[top.childIdx], top.wordIdx
            if stack[idx].childIdx += 1; stack[idx].childIdx == 26 {
                stack = stack[:idx]
            }

        }
    }

    return false
}

func (this *WordDictionary) get(letter byte) *WordDictionary {
    return this.children[letter - 'a']
}