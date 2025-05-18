package main

import (
	"errors"
	"regexp"
)

type Tokenizer struct {
	string_pattern     *regexp.Regexp
	whitespace_pattern *regexp.Regexp
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		string_pattern:     regexp.MustCompile(`\"([^"]*)\"`),
		whitespace_pattern: regexp.MustCompile(`\s+`),
	}
}

func flushCurrentToken(currentToken *string, tokens *[]string) {
	if *currentToken != "" {
		*tokens = append(*tokens, *currentToken)
		*currentToken = ""
	}
}

func tokenize(s string, t *Tokenizer) ([]string, error) {
	var tokens []string
	var currentToken string
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' || char == ')' {
			flushCurrentToken(&currentToken, &tokens)
			tokens = append(tokens, string(char))
		} else if char == '"' {
			a := t.string_pattern.FindStringIndex(s[i:])
			if a != nil && a[0] == 0 {
				tokens = append(tokens, s[i:i+a[1]])
				i += a[1] - 1
			} else {
				return nil, errors.New("unmatched quotes")
			}
		} else if t.whitespace_pattern.MatchString(string(char)) {
			flushCurrentToken(&currentToken, &tokens)
		} else {
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}
	return tokens, nil

}
