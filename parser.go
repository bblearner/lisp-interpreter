package main

import (
	"errors"
	"fmt"
	"strings"
)

type node struct {
	val        string
	childNodes []*node
	parent     *node
}

func (n node) String() string {
	if n.val == "" && len(n.childNodes) == 0 {
		return "()"
	}

	if len(n.childNodes) == 0 {
		return n.val
	} else {
		childStrs := []string{}
		for _, child := range n.childNodes {
			childStrs = append(childStrs, child.String())
		}
		return fmt.Sprintf("(%s%s)", n.val+" ", strings.Join(childStrs, " "))
	}
}

func EmptyNode() *node {
	return &node{
		val:        "",
		childNodes: nil,
		parent:     nil,
	}
}

func NewNode(val string) *node {
	return &node{
		val:        val,
		childNodes: nil,
	}
}

func parse(tokens []string) (*node, error) {
	if len(tokens) == 0 {
		return EmptyNode(), errors.New("unexpected end of input")
	}

	if tokens[0] != "(" {
		return EmptyNode(), errors.New("expected \"(\" in the beginning")
	}

	currentNode := EmptyNode()
	for _, token := range tokens[1:] {
		if token == "(" {
			temp := EmptyNode()
			temp.parent = currentNode
			currentNode.childNodes = append(currentNode.childNodes, temp)
			currentNode = temp
		} else if token == ")" {
			if currentNode.parent != nil {
				currentNode = currentNode.parent
			}
		} else {
			if currentNode.val == "" {
				currentNode.val = token
			} else {
				temp := NewNode(token)
				temp.parent = currentNode
				currentNode.childNodes = append(currentNode.childNodes, temp)
			}
		}
	}

	return currentNode, nil
}
