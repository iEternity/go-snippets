package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tree := New()
	tree.Insert("w")
	tree.Insert("wo")
	tree.Insert("wor")
	tree.Insert("word")

	if !tree.Find("word") {
		t.Fatal("word")
	}

	if !tree.Find("wo") {
		t.Fatal("wo")
	}

	if !tree.Find("wor") {
		t.Fatal("wor")
	}

	if tree.Find("wood") {
		t.Fatal("wood")
	}

	if tree.Find("words") {
		t.Fatal("words")
	}
}