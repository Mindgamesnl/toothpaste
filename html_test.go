package toothpaste

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestHelloNameHtml(t *testing.T) {
	var content = readHtmlTest("testdata/hello-name.html")

	var context = NewRenderContext()
	var renderer = NewRenderer()

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