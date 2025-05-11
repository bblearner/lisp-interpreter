package main

import "fmt"

// Add shell support

func main() {
	x := "Hello, World!"
	tokenize(x)
	fmt.Println("Hello, World!")
}

// update the tokenize funcation
// 1. We want to find a better way for reading characters
// 2. We want to add support for nested s-strings

func tokenize(s string) []string {
	var tokens []string
	var currentToken string
	for _, char := range s {
		if char == ' ' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
		} else {
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}
	return tokens

}
