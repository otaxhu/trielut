package trielut

// Trie, data structure designed to push strings of data and search for them very quickly
//
// A zero-value of this structure is ready to use methods on it.
type Trie struct {
	// LookUp Table (LUT) that stores all of the Trie's childrens (length is 256 or 0 if nil)
	lut []*Trie

	// Flag to mark this node as the end of a string in the trie
	isEnd bool
}

// Inserts b into the trie, returns false if b was already inserted into the trie, true otherwise
func (t *Trie) Insert(b []byte) bool {
	for _, char := range b {
		if t.lut == nil {
			t.lut = make([]*Trie, 256)
		}
		if t.lut[char] == nil {
			t.lut[char] = &Trie{}
		}
		t = t.lut[char]
	}

	if !t.isEnd {
		t.isEnd = true
		return true
	}

	return false
}

// Checks if b is inside of the trie, returns true if b is inside, false otherwise
func (t *Trie) Has(b []byte) bool {
	for _, char := range b {
		if t.lut == nil || t.lut[char] == nil {
			return false
		}
		t = t.lut[char]
	}

	return t.isEnd
}

// Deletes b from the trie, returns true if b was successfully deleted, false otherwise
func (t *Trie) Delete(b []byte) bool {
	var visitedNodes []*Trie

	for _, char := range b {
		if t.lut == nil || t.lut[char] == nil {
			return false
		}
		visitedNodes = append(visitedNodes, t)
		t = t.lut[char]
	}

	if !t.isEnd {
		return false
	}
	t.isEnd = false

	for i := len(b) - 1; i >= 0; i-- {
		t = visitedNodes[i]
		char := b[i]
		if len(t.lut[char].lut) > 0 || t.lut[char].isEnd {
			break
		}
		t.lut[char] = nil
	}

	return true
}
