package main

import (
	"fmt"
	"testing"
)

var parserTests = []struct {
	tokens   []string
	expected string
}{
	{
		tokens:   []string{"(", "format", "t", "\"Hello, World!\"", ")"},
		expected: `(format t "Hello, World!")`,
	},
	{
		tokens:   []string{"(", "let", "(", "(", "str", "\"Hello, world!\"", ")", ")", "(", "string-upcase", "str", ")", ")"},
		expected: `(let ( (str "Hello, world!")) (string-upcase str))`,
	},
	{
		tokens:   []string{"(", "defun", "sum", "(", "&rest", "numbers", ")", "(", "reduce", "#'+", "numbers", ")", ")"},
		expected: `(defun sum (&rest numbers) (reduce #'+ numbers))`,
	},
	{
		tokens:   []string{"(", "defmacro", "when", "(", "condition", "&body", "body", ")", "`", "(", "if", "condition", "(", "progn", "@body", ")", ")", ")", ")"},
		expected: "(defmacro when (condition &body body) ` (if condition (progn @body)))",
	},
	{
		tokens:   []string{"(", "defun", "fib", "(", "n", ")", "\"Return the nth Fibonacci number.\"", "(", "if", "(", "<", "n", "2", ")", "n", "(", "+", "(", "fib", "(", "-", "n", "1", ")", ")", "(", "fib", "(", "-", "n", "2", ")", ")", ")", ")", ")"},
		expected: `(defun fib n "Return the nth Fibonacci number." (if (< n 2) n (+ (fib (- n 1)) (fib (- n 2)))))`,
	},
	{
		tokens:   []string{"(", "defun", "fact", "(", "n", ")", "(", "if", "(", "<=", "n", "1", ")", "1", "(", "*", "n", "(", "fact", "(", "-", "n", "1", ")", ")", ")", ")", ")"},
		expected: `(defun fact n (if (<= n 1) 1 (* n (fact (- n 1)))))`,
	},
}

func TestParseNode(t *testing.T) {

	for i, test := range parserTests {

		testname := fmt.Sprintf("test # %d", i+1)
		t.Run(testname, func(t *testing.T) {
			ans, err := parse(test.tokens)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if ans.String() != test.expected {
				t.Errorf("got %v, want %v", ans, test.expected)
			}
		})
	}

}

func BenchmarkParser(b *testing.B) {
	for b.Loop() {
		for _, test := range parserTests {
			parse(test.tokens)
		}
	}
}
