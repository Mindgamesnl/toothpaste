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
	assert.Contains(t, renderer.Render(context, content), "awesome")

	context.SetVariable("cool_level", "neat")
	assert.Contains(t, renderer.Render(context, content), "neat")
}

func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}