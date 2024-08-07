package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	textContent, err := os.ReadFile("files/text.txt")
	if err != nil {
		log.Fatalf("Failed to read text.txt: %v", err)
	}

	fmt.Println("File content:")
	fmt.Println(string(textContent))

	words := strings.Fields(string(textContent))

	// Пошук слів, що починаються на голосну і закінчуються на приголосну
	wordPattern := `^[аеєиіїоуюяАЕЄИІЇОУЮЯ].*[бвгґджзйклмнпрстфхцчшщьБВГҐДЖЗЙКЛМНПРСТФХЦЧШЩЬ]$`
	wordRegex := regexp.MustCompile(wordPattern)

	fmt.Println("Found words starting with vowel and ending with consonant in text.txt:")
	for _, word := range words {
		if wordRegex.MatchString(word) {
			fmt.Println(word)
		}
	}

	// Пошук слів з однаковими буквами, розділеними будь-яким символом у файлі text.txt

	firstChar := `(\w)`
	letterPattern := fmt.Sprintf(`\\b%s\\w*%s\\b`, firstChar, firstChar)

	letterRegex := regexp.MustCompile(letterPattern)

	fmt.Println("Found words with identical letters separated by any character in text.txt:")
	for _, word := range words {
		if letterRegex.MatchString(word) {
			fmt.Println(word)
		}

	}
}
