# toothpaste
ToothPaste is an extremely basic Text/HTML formatting system with some nice utility features. It currently supports:
 - variables
 - escaped variables
 - (recursively) including other templates with passed variable context
 - using variables as arguments
 - (nested) if statements

```html
<body>

    {% if @name is joost %}
        {% if @cool_level is very %}
            Joost is awesome!
        {% end %}
        {% if @cool_level not very %}
            Joost is not very cool, but he is {{ @cool_level }} cool
        {% end %}
    {% end %}
    
    <h1>Hey there {{ @name }}!</h1>
    <p>Good to see you again</p>

    {{ include(logo) }}
</body>
```

and planned features include:
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
 - [Templating simple variables](examples/simple-variable/)
 - [Templating HTMl unescaped variables](examples/unescaped-variable/)
 - [Advanced includes](examples/include-example/)
