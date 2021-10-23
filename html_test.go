package toothpaste

import (
	"io/ioutil"
	"testing"
)

func TestHelloNameHtml(t *testing.T) {
	var content = readHtmlTest("testdata/hello-name.html")

	var context = NewRenderContext()
	var renderer = NewRenderer()

	context.SetVariable("name", "Mats")
	renderer.RegisterComponent("logo", readHtmlTest("testdata/logo.html"))

	t.Log(renderer.RecursiveRender(context, content, nil))
}

func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}