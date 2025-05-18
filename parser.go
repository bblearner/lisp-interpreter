package main

import (
	"errors"
	"fmt"
)

type node struct {
	val        string
	childNodes []*node
}

func (n node) String() string {

	if n.val != "" {
		return fmt.Sprintf("[%s]", n.val)
	} else {
		var childStrs []string
		for _, child := range n.childNodes {
			childStrs = append(childStrs, child.String())
		}
		return fmt.Sprintf("(%s)", childStrs)
	}
}

func NillNode() *node {
	return &node{
		val:        "",
		childNodes: nil,
	}
}

func NewNode(val string) *node {
	return &node{
		val:        val,
		childNodes: nil,
	}
}

func parse(tokens []string) (*node, error) {
	node, endPos, err := parseNode(tokens)
	if err != nil {
		return NillNode(), err
	}

	if len(node.childNodes) == 0 || endPos != len(tokens) {
		return NillNode(), errors.New("no child nodes found")
	} else {
		return node, nil
	}
}

func parseNode(tokens []string) (*node, int, error) {
	if len(tokens) == 0 {
		return NillNode(), 0, errors.New("unexpected end of input")
	}

	if tokens[0] != "(" {
		return NillNode(), 0, errors.New("expected \"(\" in the beginning")
	}

	node := NewNode("")
	var endPos int

	for i := 1; i < len(tokens); i++ {
		token := tokens[i]
		if token == ")" {
			return node, i + 1, nil
		}
		if token == "(" {
			childNode, endPos, err := parseNode(tokens[i:])
			if err != nil {
				return NillNode(), 0, err
			}
			node.childNodes = append(node.childNodes, childNode)
			i += endPos
		} else {
			node.childNodes = append(node.childNodes, NewNode(token))
		}
	}

	return node, endPos, nil
}
