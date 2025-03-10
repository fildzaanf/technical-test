package main

import (
	"fmt"
	"sort"
	"strings"
)

func replaceWord(rootWord []string, sentence string) string {
	sort.Slice(rootWord, func(i, j int) bool {
		return len(rootWord[i]) < len(rootWord[j])
	})

	words := strings.Split(sentence, " ")

	for i, word := range words {
		for _, root := range rootWord {
			if strings.HasPrefix(word, root) {
				words[i] = root
				break 
			}
		}
	}

	return strings.Join(words, " ")
}

func main() {
	rootWord1 := []string{"cat", "bat", "rat"}
	sentence1 := "the cattle was rattled by the battery"
	fmt.Println(replaceWord(rootWord1, sentence1))

	rootWord2 := []string{"dog", "car", "bike"}
	sentence2 := "the dogs were barking near the cars and bikers"
	fmt.Println(replaceWord(rootWord2, sentence2)) 
}
