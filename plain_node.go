package toothpaste

import (
	"errors"
	"html"
	"strconv"
)

type PlainNodeType int

const (
	PLAIN_NODE_TYPE_VARIABLE_RAW = iota
	PLAIN_NODE_TYPE_VARIABLE_ESCAPED
	PLAIN_NODE_TYPE_INCLUDE
	PLAIN_NODE_TYPE_INVALID
)

type PlainNode struct {
	start int
	end   int

	textCache    string
	hasTextCache bool
}

func (n *PlainNode) getText(root string) string {
	if n.hasTextCache {
		return n.textCache
	}

	n.textCache = root[(n.start + 3):(n.end - 3)]
	n.hasTextCache = true
	return n.textCache
}

func (n *PlainNode) parseType(root string) PlainNodeType {
	// attempt to guess the type
	var t = n.getText(root)
	var first = t[:1]
	if first == "@" {
		// is it escapable?
		if t[1:2] == "!" {
			return PLAIN_NODE_TYPE_VARIABLE_RAW
		}
		return PLAIN_NODE_TYPE_VARIABLE_ESCAPED
	} else if t[:8] == "include(" {
		return PLAIN_NODE_TYPE_INCLUDE
	}
	return PLAIN_NODE_TYPE_INVALID
}

func (n *PlainNode) evaluate(root string, c *RenderContext, r *Renderer) (string, error) {
	var nodeType = n.parseType(root)
	var content = n.getText(root)

	switch nodeType {
	case PLAIN_NODE_TYPE_VARIABLE_RAW:
		// try to find the raw variable
		value, found := c.variables[(content[2:])]
		if !found {
			return "", errors.New("Couldn't find a value for '" + (content[2:]) + "'")
		}
		return value.(string), nil
		break

	case PLAIN_NODE_TYPE_VARIABLE_ESCAPED:
		// Similar to a raw replacement, but escaped
		value, found := c.variables[(content[1:])]
		if !found {
			return "", errors.New("Couldn't find a value for '" + (content[1:]) + "'")
		}
		return html.EscapeString(value.(string)), nil
		break

	case PLAIN_NODE_TYPE_INCLUDE:
		var whatToInclude = content[8 : len(content)-1]

		// does it start with an @? if so, resolve it
		lookupValue, lookupFound := c.variables[(whatToInclude[1:])]
		if lookupFound {
			whatToInclude = lookupValue.(string)
		}

		value, found := r.components[whatToInclude]
		if !found {
			return "", errors.New("Couldn't find the component '" + whatToInclude + "'")
		}
		return r.Render(c, value)
	}

	return "", errors.New("ERROR! Couldn't evaluate '" + (content) + "' of type " + strconv.Itoa(int(nodeType)))
}
