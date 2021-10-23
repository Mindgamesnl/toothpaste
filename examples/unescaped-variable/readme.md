# Unescaped variable
This simple example templates index.html with a title and text.
All variables in this example use the `@!` annotation, making them raw and unescaped, allowing for raw html formatting.

# Expected output:
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My page</title>
</head>
<body>
  <h1><1>I am a header and not escaped!</h1></h1>
</body>
</html>

```