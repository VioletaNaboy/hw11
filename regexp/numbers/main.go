package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	numbersContent, err := os.ReadFile("files/numbers.txt")
	if err != nil {
		log.Fatalf("Failed to read numbers.txt: %v", err)
	}

	// Пошук телефонних номерів з 10 цифр
	plainPhonePattern := `\b\d{10}\b`
	plainPhoneRegex := regexp.MustCompile(plainPhonePattern)
	plainPhones := plainPhoneRegex.FindAllString(string(numbersContent), -1)

	fmt.Println("Found plain 10-digit phone numbers in numbers.txt:")
	for _, phone := range plainPhones {
		fmt.Println(phone)
	}

	// Пошук телефонних номерів у форматі з дужками
	formattedPhonePattern := `\(\d{3}\)\s?\d{3}-\d{4}`
	formattedPhoneRegex := regexp.MustCompile(formattedPhonePattern)
	formattedPhones := formattedPhoneRegex.FindAllString(string(numbersContent), -1)

	fmt.Println("Found formatted phone numbers in numbers.txt:")
	for _, phone := range formattedPhones {
		fmt.Println(phone)
	}
}
