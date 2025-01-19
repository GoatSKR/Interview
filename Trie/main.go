package main

import (
	"fmt"
)

func main() {
	t := NewTrie()

	words := []string{"hello", "world", "trie", "structure", "algorithm"}
	for _, word := range words {
		t.Insert(word)
	}

	fmt.Println("Search for 'hello':", t.Search("hello"))         // true
	fmt.Println("Search for 'world':", t.Search("world"))         // true
	fmt.Println("Search for 'trie':", t.Search("trie"))           // true
	fmt.Println("Search for 'structure':", t.Search("structure")) // true
	fmt.Println("Search for 'algorithm':", t.Search("algorithm")) // true
	fmt.Println("Search for 'notfound':", t.Search("notfound"))   // false

	fmt.Println("StartsWith 'str':", t.StartsWith("str"))   // true
	fmt.Println("StartsWith 'alg':", t.StartsWith("alg"))   // true
	fmt.Println("StartsWith 'wor':", t.StartsWith("wor"))   // true
	fmt.Println("StartsWith 'nope':", t.StartsWith("nope")) // false
}
