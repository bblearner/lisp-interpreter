package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestTokenize(t *testing.T) {

	var tests = []struct {
		s        string
		expected []string
	}{
		{`(format t "Hello, World!")`, []string{"(", "format", "t", "\"Hello, World!\"", ")"}},
		{`(defun hello () "Hello, Coding Challenge World")`, []string{"(", "defun", "hello", "(", ")", "\"Hello, Coding Challenge World\"", ")"}},
		{`(defun fib (n)
  			"Return the nth Fibonacci number."
  			(if (< n 2)
      			n
      			(+ (fib (- n 1))
         			(fib (- n 2)))))`, []string{"(", "defun", "fib", "(", "n", ")", "\"Return the nth Fibonacci number.\"", "(", "if", "(", "<", "n", "2", ")", "n", "+", "(", "fib", "(", "-", "n", "1", ")", "(", "fib", "(", "-", "n", "2", ")", ")", ")", ")", ")", ")", ")"}},
	}

	tokenizer := NewTokenizer()

	for i, test := range tests {

		testname := fmt.Sprintf("test # %d", i)
		t.Run(testname, func(t *testing.T) {
			ans, _ := tokenize(test.s, tokenizer)
			if !slices.Equal(ans, test.expected) {
				t.Errorf("got %v, want %v", ans, test.expected)
			}
		})
	}

}
