package main

import (
	"fmt"
	"slices"
	"testing"
)

var tokenizerTests = []struct {
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
				 (fib (- n 2)))))`, []string{"(", "defun", "fib", "(", "n", ")", "\"Return the nth Fibonacci number.\"", "(", "if", "(", "<", "n", "2", ")", "n", "(", "+", "(", "fib", "(", "-", "n", "1", ")", ")", "(", "fib", "(", "-", "n", "2", ")", ")", ")", ")", ")"}},
	{`(let ((str "Hello, world!"))
		(string-upcase str))`, []string{"(", "let", "(", "(", "str", "\"Hello, world!\"", ")", ")", "(", "string-upcase", "str", ")", ")"}},
	{
		`'(a (b (c d)) e)`,
		[]string{"'", "(", "a", "(", "b", "(", "c", "d", ")", ")", "e", ")"},
	},
	{
		`(defun sum (&rest numbers)
			(reduce #'+ numbers))`,
		[]string{"(", "defun", "sum", "(", "&rest", "numbers", ")", "(", "reduce", "#'+", "numbers", ")", ")"},
	},
	{
		`(defmacro when (condition &body body)
			` + "`" + `(if condition
				(progn @body))))`,
		[]string{"(", "defmacro", "when", "(", "condition", "&body", "body", ")", "`", "(", "if", "condition", "(", "progn", "@body", ")", ")", ")", ")"},
	},
	{
		`((lambda (x y) (+ x y)) 5 10)`,
		[]string{"(", "(", "lambda", "(", "x", "y", ")", "(", "+", "x", "y", ")", ")", "5", "10", ")"},
	},
}

func TestTokenize(t *testing.T) {
	tokenizer := NewTokenizer()

	for i, test := range tokenizerTests {

		testname := fmt.Sprintf("test # %d", i)
		t.Run(testname, func(t *testing.T) {
			ans, _ := tokenize(test.s, tokenizer)
			if !slices.Equal(ans, test.expected) {
				t.Errorf("got %v, want %v", ans, test.expected)
			}
		})
	}

}

func BenchmarkTokenizer(b *testing.B) {
	tokenizer := NewTokenizer()
	for b.Loop() {
		for _, test := range tokenizerTests {
			tokenize(test.s, tokenizer)
		}
	}
}
