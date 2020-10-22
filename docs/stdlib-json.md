# Module - "json"

```js
var json = import("json")
```

## Functions

- `decode(b string/bytes) => object`: Parses the JSON string and returns an
  object.
- `encode(o object) => bytes`: Returns the JSON string (bytes) of the object.
  Unlike Go's JSON package, this function does not HTML-escape texts, but, one
  can use `html_escape` function if needed.
- `indent(b string/bytes) => bytes`: Returns an indented form of input JSON
  bytes string.
- `html_escape(b string/bytes) => bytes`: Return an HTML-safe form of input
  JSON bytes string.

## Examples

```js
var json = import("json")

var encoded = json.encode({a: 1, b: [2, 3, 4]})  // JSON-encoded bytes string
var indentded = json.indent(encoded)             // indented form
var html_safe = json.html_escape(encoded)        // HTML escaped form

var decoded = json.decode(encoded)               // {a: 1, b: [2, 3, 4]}
```
