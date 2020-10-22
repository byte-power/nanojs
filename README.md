
# The Nanojs Language

**Nanojs is a minimal subset of javascript.**

Nanojs is **[fast](#benchmark)** and secure because it's compiled/executed as
bytecode on stack-based VM that's written in native Go.

```js
/* The Nanojs Language */
var fmt = import("fmt")

var each = function(seq, fn) {
    for (x in seq) { fn(x) }
}

var sum = function(init, seq) {
    each(seq, function(x) { init += x })
    return init
}

fmt.println(sum(0, [1, 2, 3]))   // "6"
fmt.println(sum("", [1, 2, 3]))  // "123"
```

## Features

- Simple and highly readable
  [Syntax](https://github.com/zeaphoo/nanojs/blob/master/docs/tutorial.md)
  - Dynamic typing with type coercion
  - Higher-order functions and closures
  - Immutable values
- [Securely Embeddable](https://github.com/zeaphoo/nanojs/blob/master/docs/interoperability.md)
  and [Extensible](https://github.com/zeaphoo/nanojs/blob/master/docs/objects.md)
- Compiler/runtime written in native Go _(no external deps or cgo)_
- Executable as a
  [standalone](https://github.com/zeaphoo/nanojs/blob/master/docs/nanojs-cli.md)
  language / REPL

## Quick Start

```
go get github.com/zeaphoo/nanojs/v2
```

A simple Go example code that compiles/runs Nanojs script code with some input/output values:

```golang
package main

import (
	"context"
	"fmt"

	"github.com/zeaphoo/nanojs/v2"
)

func main() {
	// Nanojs script code
	src := `
var each = function(seq, fn) {
    for (var x in seq) { fn(x) }
}

var sum = 0
var mul = 1
each([a, b, c, d], function(x) {
	sum += x
	mul *= x
})`

	// create a new Script instance
	script := nanojs.NewScript([]byte(src))

	// set values
	_ = script.Add("a", 1)
	_ = script.Add("b", 9)
	_ = script.Add("c", 8)
	_ = script.Add("d", 4)

	// run the script
	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}

	// retrieve values
	sum := compiled.Get("sum")
	mul := compiled.Get("mul")
	fmt.Println(sum, mul) // "22 288"
}
```

## References

- [Language Syntax](https://github.com/zeaphoo/nanojs/blob/master/docs/tutorial.md)
- [Object Types](https://github.com/zeaphoo/nanojs/blob/master/docs/objects.md)
- [Runtime Types](https://github.com/zeaphoo/nanojs/blob/master/docs/runtime-types.md)
  and [Operators](https://github.com/zeaphoo/nanojs/blob/master/docs/operators.md)
- [Builtin Functions](https://github.com/zeaphoo/nanojs/blob/master/docs/builtins.md)
- [Interoperability](https://github.com/zeaphoo/nanojs/blob/master/docs/interoperability.md)
- [Nanojs CLI](https://github.com/zeaphoo/nanojs/blob/master/docs/nanojs-cli.md)
- [Standard Library](https://github.com/zeaphoo/nanojs/blob/master/docs/stdlib.md)
