package toothpaste

import (
	"io/ioutil"
	"testing"
)

func TestHelloNameHtml(t *testing.T) {
	var content = readHtmlTest("testdata/hello-name.html")

	var context = NewRenderContext()
	var renderer = NewRenderer()

	context.SetVariable("name", "joost")
	context.SetVariable("cool_level", "fucking")
	renderer.RegisterComponent("logo", readHtmlTest("testdata/logo.html"))

	var body = renderer.Render(context, content)
	t.Log(body)
}

func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}