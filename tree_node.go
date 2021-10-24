package toothpaste

import (
	"errors"
	"strconv"
	"strings"
)

type TreeNodeType int

const (
	TREE_NODE_TYPE_END = iota
	TREE_NODE_TYPE_IF
	TREE_NODE_TYPE_INVALID
)

type TreeNode struct {
	start int
	end   int

	contentStart int
	contentEnd   int
	content      string

	hasTextCache bool
	parent       *TreeNode
	children     []*TreeNode
	nodeType     TreeNodeType
	tag          string
	isRoot       bool

	failure error
}

func (n *TreeNode) getText(root string) string {
	if n.hasTextCache {
		return n.tag
	}

	n.tag = root[(n.start + 3):(n.end - 3)]
	n.hasTextCache = true
	return n.tag
}

func (n *TreeNode) addChildren(tn *TreeNode) {
	n.children = append(n.children, tn)
}

func (n *TreeNode) parseType(root string) TreeNodeType {
	// attempt to guess the type
	var t = n.getText(root)

	if t == "end" {
		return TREE_NODE_TYPE_END
	} else if t[:2] == "if" {
		return TREE_NODE_TYPE_IF
	}

	return TREE_NODE_TYPE_INVALID
}

func (n *TreeNode) makeErrorNode(message string) []TreeReturnReplacement {
	n.failure = errors.New(message)

	return []TreeReturnReplacement{
		{
			start: n.contentStart,
			end:   n.contentEnd,
		},
		{
			start:    n.start,
			end:      n.end,
			newValue: message,
		},
	}
}

func (n *TreeNode) makeSelfReplacement(alsoRemoveBody bool) []TreeReturnReplacement {
	if alsoRemoveBody {
		return []TreeReturnReplacement{
			{
				start: n.start,
				end:   n.end,
			},
			{
				start: n.contentStart,
				end:   n.contentEnd,
			},
		}
	}
	return []TreeReturnReplacement{
		{
			start: n.start,
			end:   n.end,
		},
	}
}

func (n *TreeNode) render(c *RenderContext) []TreeReturnReplacement {
	if n.isRoot {
		var responses []TreeReturnReplacement
		for i := range n.children {
			var renderResults = n.children[i].render(c)
			for i2 := range renderResults {
				responses = append(responses, renderResults[i2])
			}
			if n.children[i].failure != nil {
				n.failure = n.children[i].failure
			}
		}
		return responses
	}

	if n.nodeType == TREE_NODE_TYPE_END {
		return n.makeSelfReplacement(false)
	}

	if n.nodeType == TREE_NODE_TYPE_IF {
		// if statement parsing and logic
		var parts = strings.Split(n.tag, " ")
		var lookupVariable, selector, expected = parts[1], parts[2], parts[3]
		var lookupValue, lookupFound = c.getVariable(lookupVariable[1:])

		if !lookupFound {
			// false, so end here
			return n.makeErrorNode("<b>couldn't find variable " + lookupVariable[1:] + "</b>")
		}

		// handle statement
		if selector == "is" {
			if lookupValue != expected {
				return n.makeSelfReplacement(true)
			}
		} else if selector == "not" {
			if lookupValue == expected {
				return n.makeSelfReplacement(true)
			}
		} else {
			return n.makeErrorNode("<b>unknown state " + selector + "</b>")
		}

		// render children
		if len(n.children) == 0 {
			return n.makeSelfReplacement(false)
		} else {
			var responses = n.makeSelfReplacement(false)
			for i := range n.children {
				var renderResults = n.children[i].render(c)
				for i2 := range renderResults {
					responses = append(responses, renderResults[i2])
				}

				// does it have an error?
				if n.children[i].failure != nil {
					n.failure = n.children[i].failure
				}

			}
			return responses
		}
	}

	return n.makeErrorNode("<b>unknown tag type " + strconv.Itoa(int(n.nodeType)) + "</b>")
}
