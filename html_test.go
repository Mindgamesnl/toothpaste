package toothpaste

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestVariableSetter(t *testing.T) {
	var renderer = NewRenderer()
	var context = NewRenderContext()

	value, err := renderer.Render(context, readHtmlTest("testdata/set.html"))

	assert.Nil(t, err)
	assert.Contains(t, value, "value!")
}

func TestHelloNameHtml(t *testing.T) {
	var content = readHtmlTest("testdata/hello-name.html")

	var renderer = NewRenderer()
	var context = NewRenderContext()
	context.SetVariable("user_state", func(ctx *RenderContext) string {
		// check if the user is logged in, and return the state
		return "logged_in"
	})

	context.SetVariable("name", "Mats")

	context.SetVariable("name", "joost")
	renderer.RegisterComponent("logo", readHtmlTest("testdata/logo.html"))

	context.SetVariable("cool_level", "very")
	renderOutput, err := renderer.Render(context, content)
	assert.Nil(t, err)
	assert.Contains(t, renderOutput, "awesome")

	context.SetVariable("cool_level", "neat")
	renderOutput, err = renderer.Render(context, content)
	assert.Nil(t, err)
	assert.Contains(t, renderOutput, "neat")
}

func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}
