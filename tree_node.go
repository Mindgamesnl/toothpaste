package toothpaste

import (
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

func (n *TreeNode) render(c *RenderContext) []TreeReturnReplacement {
	if n.isRoot {
		var responses []TreeReturnReplacement
		for i := range n.children {
			var renderResults = n.children[i].render(c)
			for i2 := range renderResults {
				responses = append(responses, renderResults[i2])
			}
		}
		return responses
	}

	if n.nodeType == TREE_NODE_TYPE_END {
		return []TreeReturnReplacement{
			TreeReturnReplacement{
				start: n.start,
				end: n.end,
				newValue: "",
			},
		}
	}

	if n.nodeType == TREE_NODE_TYPE_IF {
		// if statement parsing and logic
		var parts = strings.Split(n.tag, " ")
		var lookupVariable, selector, expected = parts[1], parts[2], parts[3]
		var lookupValue, lookupFound = c.getVariable(lookupVariable[1:])

		if !lookupFound {
			// false, so end here
			return []TreeReturnReplacement{
				TreeReturnReplacement{
					start: n.start,
					end: n.end,
					newValue: "<b>Error! couldn't find variable " + lookupVariable[1:] + "</b>",
				},
			}
		}

		if selector == "is" {
			if lookupValue != expected {
				return []TreeReturnReplacement{
					TreeReturnReplacement{
						start: n.start,
						end: n.end,
						newValue: "",
					},
					TreeReturnReplacement{
						start: n.contentStart,
						end: n.contentEnd,
						newValue: "",
					},
				}
			}
		} else if selector == "not" {
			if lookupValue == expected {
				return []TreeReturnReplacement{
					TreeReturnReplacement{
						start: n.start,
						end: n.end,
						newValue: "",
					},
					TreeReturnReplacement{
						start: n.contentStart,
						end: n.contentEnd,
						newValue: "",
					},
				}
			}
		} else {
			return []TreeReturnReplacement{
				TreeReturnReplacement{
					start: n.start,
					end: n.end,
					newValue: "Unknown state " + selector,
				},
			}
		}

		if len(n.children) == 0 {
			return []TreeReturnReplacement{
				TreeReturnReplacement{
					start: n.start,
					end: n.end,
					newValue: "",
				},
			}
		} else {
			var responses = []TreeReturnReplacement{
				TreeReturnReplacement{
					start: n.start,
					end: n.end,
					newValue: "",
				},
			}
			for i := range n.children {
				var renderResults = n.children[i].render(c)
				for i2 := range renderResults {
					responses = append(responses, renderResults[i2])
				}
			}
			return responses
		}
	}

	return []TreeReturnReplacement{
		TreeReturnReplacement{
			start: n.start,
			end: n.end,
			newValue: "Unknown type " + strconv.Itoa(int(n.nodeType)),
		},
	}
}
