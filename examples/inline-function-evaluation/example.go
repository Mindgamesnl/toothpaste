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
	context.SetVariable("name", "Mindgamesnl")

	// set a function to check login state, you'd run your session checks in here
	context.SetVariable("user_state", func(ctx *toothpaste.RenderContext) string {
		return "logged_in"
	})

	// render
	fmt.Println(renderer.Render(context, readHtmlTest("index.html")))
}


// util
func readHtmlTest(f string) string {
	content, _ := ioutil.ReadFile(f)
	return string(content)
}

