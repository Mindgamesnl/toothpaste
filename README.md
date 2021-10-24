# toothpaste
ToothPaste is an extremely basic Text/HTML formatting system with some nice utility features. It currently supports:
 - variables
 - escaped variables
 - (recursively) including other templates with passed variable context
 - using variables as arguments
 - (nested) if statements
 - template functions

## Variables
Variables can be a `string`, `int`, `int8`, `int16`, `int64` and `func(ctx *RenderContext) string`. Variable types (or processors) will be resolved everytime a variable is invoked from your template. You can use functions to process scripts during rendering.

example:
```go
var context = NewRenderContext()
context.SetVariable("user_state", func(ctx *RenderContext) string {
    // check if the user is logged in, and return the state
    return "logged_in"
})

context.SetVariable("name", "Mats")
```
```html
{{ include(navbar) }}


{% if @user_state is logged_in %}
    <h1>You are logged in! welcome back, {{ @name }}!</h1>
{% end %}
{% if @user_state not logged_in %}
    <h1>You seem to be new! please <a href="/login">login</a>.</h1>
    <p>your state is {{ @user_state }}</p>
{% end %}

{{ include(footer) }}
```

## Notes
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
