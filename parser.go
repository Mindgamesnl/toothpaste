package toothpaste

import "regexp"

var SEARCH_PATTERN *regexp.Regexp

func getPatternMatcher() *regexp.Regexp {
	if SEARCH_PATTERN == nil {
		r, _ := regexp.Compile(`{{ [^/{}]* }}`)
		SEARCH_PATTERN = r
	}
	return SEARCH_PATTERN
}

func findNodesIn(root string) []Node {
	var nodes []Node
	var elements = getPatternMatcher().FindAllStringIndex(root, -1)
	for i := range elements {
		var element = elements[i]

		nodes = append(nodes, Node{
			start: element[0],
			end: element[1],
			hasTextCache: false,
		})
	}
	return nodes
}
