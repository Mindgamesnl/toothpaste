# Include variable
Toothpaste supports simple including of components. Components need to be registered by a namespace, and can then be used as `{{ include(navbar) }}` if you have a component registered by that name. Note that you are allowed to use variables in this case, so `{{ include(@page) }}` is a perfectly valid syntax as long as page is a devined variable in the context, and that variable is the name of a valid component

# Expected output:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My page</title>
</head>
<body>
<div style="text-align: center;">
    <h1>My Awesome Website!</h1>
    <hr>
</div>
<h2>Welcome Mindgamesnl!</h2>
</body>
</html>

```