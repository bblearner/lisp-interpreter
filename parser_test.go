package main

import (
	"fmt"
	"testing"
)

func TestParseNode(t *testing.T) {

	var tests = []struct {
		tokens   []string
		expected *node
	}{
		{
			tokens:   []string{"(", "format", "t", "\"Hello, World!\"", ")"},
			expected: &node{val: "", childNodes: []*node{{val: "format", childNodes: nil}, {val: "t", childNodes: nil}, {val: "\"Hello, World!\"", childNodes: nil}}},
		},
		{
			tokens:   []string{"(", "let", "(", "(", "str", "\"Hello, world!\"", ")", ")", "(", "string-upcase", "str", ")", ")"},
			expected: &node{val: "", childNodes: []*node{{val: "let", childNodes: nil}, {val: "t", childNodes: nil}, {val: "\"Hello, World!\"", childNodes: nil}}},
		},
	}

	for i, test := range tests {

		testname := fmt.Sprintf("test # %d", i)
		t.Run(testname, func(t *testing.T) {
			ans, err := parse(test.tokens)

			fmt.Println("ans:", ans)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if ans.val != test.expected.val || len(ans.childNodes) != len(test.expected.childNodes) {
				t.Errorf("got %v, want %v", ans, test.expected)
			}
			for j := range ans.childNodes {
				if ans.childNodes[j].val != test.expected.childNodes[j].val {
					t.Errorf("got %v, want %v", ans.childNodes[j], test.expected.childNodes[j])
				}
			}
		})
	}

}
