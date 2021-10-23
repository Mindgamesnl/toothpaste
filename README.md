# toothpaste
ToothPaste is an extremely basic Text/HTML formatting system with some nice utility features. It currently supports:
 - variables
 - escaped variables
 - (recursively) including other templates with passed variable context
 - using variables as arguments

and planned features include:
 - IF statements
 - For loops

## What to expect
```go
// setup a renderer
var renderer = toothpaste.NewRenderer()

// register global component in the renderer
renderer.RegisterComponent("header", readHtmlTest("header.html"))

// setup a context for this specific render
var context = toothpaste.NewRenderContext()

// set variables
context.SetVariable("title", "My page")
context.SetVariable("name", "Mindgamesnl")

// render
fmt.Println(renderer.Render(context, readHtmlTest("index.html")))
```

## Examples
 - [Templating simple variables](examples/simple-variable/readme.md)
 - [Templating HTMl unescaped variables](examples/unescaped-variable/readme.md)
 - [Advanced includes](examples/include-example/readme.md)