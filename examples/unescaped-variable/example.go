package main

import (
	"fmt"
	"github.com/Mindgamesnl/toothpaste"
	"io/ioutil"
)

func main() {
	// setup a renderer
	var renderer = toothpaste.NewRenderer()

	// setup a context for this specific render
	var context = toothpaste.NewRenderContext()

	// set variables
	context.SetVariable("title", "My page")
	context.SetVariable("text", "<1>I am a header and not escaped!</h1>")

	// render
	fmt.Println(renderer.Render(context, readHtmlTest("index.html")))
}


// util
func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}

