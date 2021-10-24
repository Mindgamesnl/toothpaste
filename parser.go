package toothpaste

import "regexp"

var PLAIN_SEARCH_PATTERN *regexp.Regexp
var TREE_SEARCH_PATTERN *regexp.Regexp
var CLEANER *regexp.Regexp

//

func getCleanerPattern() *regexp.Regexp {
	if CLEANER == nil {
		r, _ := regexp.Compile(`[ \t]`)
		CLEANER = r
	}
	return CLEANER
}

func getTreePatternMatcher() *regexp.Regexp {
	if TREE_SEARCH_PATTERN == nil {
		r, _ := regexp.Compile(`{\% [^/{}]* \%}`)
		TREE_SEARCH_PATTERN = r
	}
	return TREE_SEARCH_PATTERN
}

func getPlainpatternMatcher() *regexp.Regexp {
	if PLAIN_SEARCH_PATTERN == nil {
		r, _ := regexp.Compile(`{{ [^/{}]* }}`)
		PLAIN_SEARCH_PATTERN = r
	}
	return PLAIN_SEARCH_PATTERN
}

func findTreeNodes(root string) []TreeNode {
	var nodes []TreeNode
	var elements = getTreePatternMatcher().FindAllStringIndex(root, -1)
	for i := range elements {
		var element = elements[i]

		nodes = append(nodes, TreeNode{
			start: element[0],
			end: element[1],
			hasTextCache: false,
		})
	}
	return nodes
}

func findPlainNodes(root string) []PlainNode {
	var nodes []PlainNode
	var elements = getPlainpatternMatcher().FindAllStringIndex(root, -1)
	for i := range elements {
		var element = elements[i]

		nodes = append(nodes, PlainNode{
			start: element[0],
			end: element[1],
			hasTextCache: false,
		})
	}
	return nodes
}
