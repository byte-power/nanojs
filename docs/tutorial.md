# Nanojs Language Syntax

Nanojs's syntax is designed to be familiar to Go developers while being a bit
simpler and more streamlined.


## Values and Value Types

In Nanojs, everything is a value, and, all values are associated with a type.

```golang
19 + 84               // int values
"aomame" + `kawa`     // string values
-9.22 + 1e10          // float values
true || false         // bool values
'九' > '9'             // char values
[1, false, "foo"]     // array value
{a: 12.34, b: "bar"}  // map value
func() { /*...*/ }    // function value
```

Here's a list of all available value types in Nanojs.

| Nanojs Type | Description | Equivalent Type in Go |
| :---: | :---: | :---: |
| int | signed 64-bit integer value | `int64` |
| float | 64-bit floating point value | `float64` |
| bool | boolean value | `bool` |
| char | unicode character | `rune` |
| string | unicode string | `string` |
| bytes | byte array | `[]byte` |
| error | [error](#error-values) value | - |
| time | time value | `time.Time` |
| array | value array _(mutable)_ | `[]interface{}` |
| immutable array | [immutable](#immutable-values) array | - |
| map | value map with string keys _(mutable)_ | `map[string]interface{}` |
| immutable map | [immutable](#immutable-values) map | - |
| undefined | [undefined](#undefined-values) value | - |
| function | [function](#function-values) value | - |
| _user-defined_ | value of [user-defined types](https://github.com/zeaphoo/nanojs/blob/master/docs/objects.md) | - |

### Error Values

In Nanojs, an error can be represented using "error" typed values. An error
value is created using `error` expression, and, it must have an underlying
value. The underlying value of an error value can be access using `.value`
selector.

```js
var err1 = error("oops")    // error with string value
var err2 = error(1+2+3)     // error with int value
if (is_error(err1)) {      // 'is_error' builtin function
  var err_val = err1.value  // get underlying value
}
```

### Immutable Values

In Nanojs, basically all values (except for array and map) are immutable.

```js
var s = "12345"
s[1] = 'b'    // illegal: String is immutable

var a = [1, 2, 3]
a[1] = "two"  // ok: a is now [1, "two", 3]
```

An array or map value can be made immutable using `immutable` expression.

```js
var b = immutable([1, 2, 3])
b[1] = "foo"  // illegal: 'b' references to an immutable array.
```

Note that re-assigning a new value to the variable has nothing to do with the
value immutability.

```js
var s = "abc"
s = "foo"                  // ok
var a = immutable([1, 2, 3])
a = false                  // ok
```

Note that, if you copy (using `copy` builtin function) an immutable value, it
will return a "mutable" copy. Also, immutability is not applied to the
individual elements of the array or map value, unless they are explicitly made
immutable.

```js
var a = immutable({b: 4, c: [1, 2, 3]})
a.b = 5        // illegal
a.c[1] = 5     // ok: because 'a.c' is not immutable

a = immutable({b: 4, c: immutable([1, 2, 3])})
a.c[1] = 5     // illegal
```

### Undefined Values

In Nanojs, an "undefined" value can be used to represent an unexpected or
non-existing value:

- A function that does not return a value explicitly considered to return
`undefined` value.
- Indexer or selector on composite value types may return `undefined` if the
key or index does not exist.
- Type conversion builtin functions without a default value will return
`undefined` if conversion fails.

```js
var a = function() { var b = 4 }()    // a == undefined
var b = [1, 2, 3][10]          // b == undefined
var c = {a: "foo"}["b"]        // c == undefined
var d = int("foo")             // d == undefined
```

### Array Values

In Nanojs, array is an ordered list of values of any types. Elements of an array
can be accessed using indexer `[]`.

```js
[1, 2, 3][0]       // == 1
[1, 2, 3][2]       // == 3
[1, 2, 3][3]       // == undefined

["foo", "bar", [1, 2, 3]]   // ok: array with an array element
```

### Map Values

In Nanojs, map is a set of key-value pairs where key is string and the value is
of any value types. Value of a map can be accessed using indexer `[]` or
selector '.' operators.

```js
var m = { a: 1, b: false, c: "foo" }
m["b"]                                // == false
m.c                                   // == "foo"
m.x                                   // == undefined

{a: [1,2,3], b: {c: "foo", d: "bar"}} // ok: map with an array element and a map element
```

### Function Values

In Nanojs, function is a callable value with a number of function arguments and
a return value. Just like any other values, functions can be passed into or
returned from another function.

```js
var my_func = function(arg1, arg2) {
  return arg1 + arg2
}

var adder = function(base) {
  return func(x) { return base + x }  // capturing 'base'
}
var add5 = adder(5)
var nine = add5(4)    // == 9
```

Unlike Go, Nanojs does not have declarations. So the following code is illegal:

```js
function my_func(arg1, arg2) {  // illegal
  return arg1 + arg2
}
```

Nanojs also supports variadic functions/closures:

```js
var variadic = function(a, b, ...c) {
  return [a, b, c]
}
variadic(1, 2, 3, 4) // [1, 2, [3, 4]]

var variadicClosure = function(a) {
  return function(b, ...c) {
    return [a, b, c]
  }
}
variadicClosure(1)(2, 3, 4) // [1, 2, [3, 4]]
```

Only the last parameter can be variadic. The following code is also illegal:

```js
// illegal, because a is variadic and is not the last parameter
var illegal = function(a..., b) { /*... */ }
```

When calling a function, the number of passing arguments must match that of
function definition.

```js
var f = function(a, b) {}
f(1, 2, 3) // Runtime Error: wrong number of arguments: want=2, got=3
```

Like Go, you can use ellipsis `...` to pass array-type value as its last parameter:

```js
var f1 = function(a, b, c) { return a + b + c }
f1([1, 2, 3]...)    // => 6
f1(1, [2, 3]...)    // => 6
f1(1, 2, [3]...)    // => 6
f1([1, 2]...)       // Runtime Error: wrong number of arguments: want=3, got=2

var f2 = function(a, ...b) {}
f2(1)               // valid; a = 1, b = []
f2(1, 2)            // valid; a = 1, b = [2]
f2(1, 2, 3)         // valid; a = 1, b = [2, 3]
f2([1, 2, 3]...)    // valid; a = 1, b = [2, 3]
```

## Variables and Scopes

A value can be assigned to a variable using assignment operator `:=` and `=`.

- `:=` operator defines a new variable in the scope and assigns a value.
- `=` operator assigns a new value to an existing variable in the scope.

Variables are defined either in global scope (defined outside function) or in
local scope (defined inside function).

```js
var a = "foo"      // define 'a' in global scope

function() {        // function scope A
  var b = 52       // define 'b' in function scope A

  function() {      // function scope B
    var c = 19.84  // define 'c' in function scope B

    a = "bee"   // ok: assign new value to 'a' from global scope
    b = 20      // ok: assign new value to 'b' from function scope A

    b = true   // ok: define new 'b' in function scope B
                //     (shadowing 'b' from function scope A)
  }

  a = "bar"     // ok: assigne new value to 'a' from global scope
  b = 10        // ok: assigne new value to 'b'
  var a = -100     // ok: define new 'a' in function scope A
                //     (shadowing 'a' from global scope)

  c = -9.1      // illegal: 'c' is not defined
  var b = [1, 2]   // illegal: 'b' is already defined in the same scope
}

b = 25          // illegal: 'b' is not defined
var a = {d: 2}     // illegal: 'a' is already defined in the same scope
```

Unlike Go, a variable can be assigned a value of different types.

```js
var a = 123        // assigned    'int'
a = "123"       // re-assigned 'string'
a = [1, 2, 3]   // re-assigned 'array'
```

## Type Conversions

Although the type is not directly specified in Nanojs, one can use type
conversion
[builtin functions](https://github.com/zeaphoo/nanojs/blob/master/docs/builtins.md)
to convert between value types.

```js
s1 := string(1984)    // "1984"
i2 := int("-999")     // -999
f3 := float(-51)      // -51.0
b4 := bool(1)         // true
c5 := char("X")       // 'X'
```

See [Operators](https://github.com/zeaphoo/nanojs/blob/master/docs/operators.md)
for more details on type coercions.

## Operators

### Unary Operators

| Operator | Usage | Types |
| :---: | :---: | :---: |
| `+`   | same as `0 + x` | int, float |
| `-`   | same as `0 - x` | int, float |
| `!`   | logical NOT | all types* |
| `^`   | bitwise complement | int |

_In Nanojs, all values can be either truthy or falsy._

### Binary Operators

| Operator | Usage | Types |
| :---: | :---: | :---: |
| `==` | equal | all types |
| `!=` | not equal | all types |
| `&&` | logical AND | all types |
| `\|\|` | logical OR | all types |
| `+`   | add/concat | int, float, string, char, time, array |
| `-`   | subtract | int, float, char, time |
| `*`   | multiply | int, float |
| `/`   | divide | int, float |
| `&`   | bitwise AND | int |
| `\|`   | bitwise OR | int |
| `^`   | bitwise XOR | int |
| `&^`   | bitclear (AND NOT) | int |
| `<<`   | shift left | int |
| `>>`   | shift right | int |
| `<`   | less than | int, float, char, time, string |
| `<=`   | less than or equal to | int, float, char, time, string |
| `>`   | greater than | int, float, char, time, string |
| `>=`   | greater than or equal to | int, float, char, time, string |

_See [Operators](https://github.com/zeaphoo/nanojs/blob/master/docs/operators.md)
for more details._

### Ternary Operators

Nanojs has a ternary conditional operator `(condition expression) ? (true expression) : (false expression)`.

```js
var a = true ? 1 : -1    // a == 1

var min = function(a, b) {
  return a < b ? a : b
}
var b = min(5, 10)      // b == 5
```

### Assignment and Increment Operators

| Operator | Usage |
| :---: | :---: |
| `+=` | `(lhs) = (lhs) + (rhs)` |
| `-=` | `(lhs) = (lhs) - (rhs)` |
| `*=` | `(lhs) = (lhs) * (rhs)` |
| `/=` | `(lhs) = (lhs) / (rhs)` |
| `%=` | `(lhs) = (lhs) % (rhs)` |
| `&=` | `(lhs) = (lhs) & (rhs)` |
| `\|=` | `(lhs) = (lhs) \| (rhs)` |
| `&^=` | `(lhs) = (lhs) &^ (rhs)` |
| `^=` | `(lhs) = (lhs) ^ (rhs)` |
| `<<=` | `(lhs) = (lhs) << (rhs)` |
| `>>=` | `(lhs) = (lhs) >> (rhs)` |
| `++` | `(lhs) = (lhs) + 1` |
| `--` | `(lhs) = (lhs) - 1` |

### Operator Precedences

Unary operators have the highest precedence, and, ternary operator has the
lowest precedence. There are five precedence levels for binary operators.
Multiplication operators bind strongest, followed by addition operators,
comparison operators, `&&` (logical AND), and finally `||` (logical OR):

| Precedence | Operator |
| :---: | :---: |
| 5 | `*`  `/`  `%`  `<<`  `>>`  `&`  `&^` |
| 4 | `+`  `-`  `\|`  `^` |
| 3 | `==`  `!=`  `<`  `<=`  `>`  `>=` |
| 2 | `&&` |
| 1 | `\|\|` |

Like Go, `++` and `--` operators form statements, not expressions, they fall
outside the operator hierarchy.

### Selector and Indexer

One can use selector (`.`) and indexer (`[]`) operators to read or write
elements of composite types (array, map, string, bytes).

```js
["one", "two", "three"][1]  // == "two"

var m = {
  a: 1,
  b: [2, 3, 4],
  c: func() { return 10 }
}
m.a              // == 1
m["b"][1]        // == 3
m.c()            // == 10
m.x = 5          // add 'x' to map 'm'
m["b"][5]        // == undefined
m["b"][5].d      // == undefined
m.b[5] = 0       // == undefined
m.x.y.z          // == undefined
```

Like Go, one can use slice operator `[:]` for sequence value types such as
array, string, bytes.

```js
var a = [1, 2, 3, 4, 5][1:3]    // == [2, 3]
var b = [1, 2, 3, 4, 5][3:]     // == [4, 5]
var c = [1, 2, 3, 4, 5][:3]     // == [1, 2, 3]
var d = "hello world"[2:10]     // == "llo worl"
var c = [1, 2, 3, 4, 5][-1:10]  // == [1, 2, 3, 4, 5]
```

**Note: Keywords cannot be used as selectors.**

```js
var a = {in: true} // Parse Error: expected map key, found 'in'
a.function = ""     // Parse Error: expected selector, found 'function'
```

Use double quotes and indexer to use keywords with maps.

```js
var a = {"in": true}
a["function"] = ""
```

## Statements

### If Statement

"If" statement is very similar to Go.

```js
if (a < 0) {
  // execute if 'a' is negative
} else if (a == 0) {
  // execute if 'a' is zero
} else {
  // execute if 'a' is positive
}
```

Like Go, the condition expression may be preceded by a simple statement,
which executes before the expression is evaluated.

```js
if (var a = foo(); a < 0) {
  // execute if 'a' is negative
}
```

### For Statement

"For" statement is very similar to Go.

```js
// for (init); (condition); (post) {}
for (a:=0; a<10; a++) {
  // ...
}

// for (condition) {}
for (a < 10) {
  // ...
}

// for {}
for {
  // ...
}
```

### For-In Statement

"For-In" statement is new in Nanojs. It's similar to Go's `for range` statement.
"For-In" statement can iterate any iterable value types (array, map, bytes,
string, undefined).

```js
for (var v in [1, 2, 3]) {          // array: element
  // 'v' is value
}
for (var i, v in [1, 2, 3]) {       // array: index and element
  // 'i' is index
  // 'v' is value
}
for (var k, v in {k1: 1, k2: 2}) {  // map: key and value
  // 'k' is key
  // 'v' is value
}
```

## Modules

Module is the basic compilation unit in Nanojs. A module can import another
module using `import` expression.

Main module:

```js
var sum = import("./sum")  // load module from a local file
fmt.print(sum(10))      // module function
```

Another module in `sum.js` file:

```js
var base = 5

export function(x) {
  return x + base
}
```

In Nanojs, modules are very similar to functions.

- `import` expression loads the module code and execute it like a function.
- Module should return a value using `export` statement.
  - Module can return `export` a value of any types: int, map, function, etc.
  - `export` in a module is like `return` in a function: it stops execution and
  return a value to the importing code.
  - `export`-ed values are always immutable.
  - If the module does not have any `export` statement, `import` expression
  simply returns `undefined`. _(Just like the function that has no `return`.)_
  - Note that `export` statement is completely ignored and not evaluated if
  the code is executed as a main module.

Also, you can use `import` expression to load the
[Standard Library](https://github.com/zeaphoo/nanojs/blob/master/docs/stdlib.md) as
well.

```js
var math = import("math")
var a = math.abs(-19.84)  // == 19.84
```

## Comments

Like Go, Nanojs supports line comments (`//...`) and block comments
(`/* ... */`).

```js
/*
  multi-line block comments
*/

var a = 5    // line comments
```