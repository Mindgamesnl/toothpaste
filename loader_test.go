package toothpaste

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToothpasteLoader(t *testing.T)  {
	var elements = findNodesIn(`Good to see you, {{ name }}! how are you? and how is your son {{ son_name }}`)
	assert.Len(t, elements, 2)
}

func TestTypeGuesser(t *testing.T)  {
	var root = `Good to see you, {{ @name }}! how are you? and how is your son {{ @!son_name }}`
	var elements = findNodesIn(root)
	for i := range elements {
		assert.NotEqual(t, NODE_TYPE_INVALID, elements[i].parseType(root))
	}
}

func TestNodeEvaluator(t *testing.T)  {
	var root = `Good to see you, {{ @name }}! how are you? and how is your son {{ @!son_name }}`

	var context = NewRenderContext()
	// evil value
	context.SetVariable("name", "<evil>bart</evil>")
	// raw value
	context.SetVariable("son_name", "<h1>joost</h1>")

	var elements = findNodesIn(root)

	for i := range elements {
		var value, err = elements[i].evaluate(root, context, NewRenderer(), nil)
		assert.Nil(t, err)

		if strings.Contains(value, "bart") {
			assert.NotContainsf(t, value, "<", "Fail bart")
		} else {
			assert.Containsf(t, value, "<", "Fail joost")
		}
	}
}

func TestRender(t *testing.T)  {
	var root = `Good to see you, {{ @name }}! how are you? and how is your son {{ @!son_name }}?`

	var context = NewRenderContext()
	var renderer = NewRenderer()

	// evil value
	context.SetVariable("name", "Bart")
	// raw value
	context.SetVariable("son_name", "Joost")

	// check if the output is formatted perfectly
	assert.Equal(t, renderer.RecursiveRender(context, root, nil), "Good to see you, Bart! how are you? and how is your son Joost?")
}


