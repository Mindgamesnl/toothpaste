# toothpaste
Toothpaste is a simple templating engine for Go web applications, inspired by Blade/Twig.
It supports basic if-statements, variables, inline variable definitions, variable HTML sanitization by default (but also allows raw printing), recursive template importing with context and processing functions. 

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

{% set app_name "My cool app!" %}

{% if @user_state is logged_in %}
    <h1>You are logged in! welcome back to {{ @app_name }}, {{ @name }}!</h1>
{% end %}
{% if @user_state not logged_in %}
    <h1>You seem to be new! please <a href="/login">login</a>.</h1>
    <p>your state is {{ @user_state }}</p>
{% end %}

{{ include(footer) }}
```

## Examples
 - [Login check with functional variables](examples/inline-function-evaluation/)
 - [Templating simple variables](examples/simple-variable/)
 - [Templating HTMl unescaped variables](examples/unescaped-variable/)
 - [Advanced includes](examples/include-example/)
