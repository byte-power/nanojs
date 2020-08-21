package nanojs_test

import (
	"testing"

	"github.com/zeaphoo/nanojs/v2"
	"github.com/zeaphoo/nanojs/v2/require"

)

func TestObject_TypeName(t *testing.T) {
	var o nanojs.Object = &nanojs.Int{}
	require.Equal(t, "int", o.TypeName())
	o = &nanojs.Float{}
	require.Equal(t, "float", o.TypeName())
	o = &nanojs.Char{}
	require.Equal(t, "char", o.TypeName())
	o = &nanojs.String{}
	require.Equal(t, "string", o.TypeName())
	o = &nanojs.Bool{}
	require.Equal(t, "bool", o.TypeName())
	o = &nanojs.Array{}
	require.Equal(t, "array", o.TypeName())
	o = &nanojs.Map{}
	require.Equal(t, "map", o.TypeName())
	o = &nanojs.ArrayIterator{}
	require.Equal(t, "array-iterator", o.TypeName())
	o = &nanojs.StringIterator{}
	require.Equal(t, "string-iterator", o.TypeName())
	o = &nanojs.MapIterator{}
	require.Equal(t, "map-iterator", o.TypeName())
	o = &nanojs.BuiltinFunction{Name: "fn"}
	require.Equal(t, "builtin-function:fn", o.TypeName())
	o = &nanojs.UserFunction{Name: "fn"}
	require.Equal(t, "user-function:fn", o.TypeName())
	o = &nanojs.CompiledFunction{}
	require.Equal(t, "compiled-function", o.TypeName())
	o = &nanojs.Undefined{}
	require.Equal(t, "undefined", o.TypeName())
	o = &nanojs.Error{}
	require.Equal(t, "error", o.TypeName())
	o = &nanojs.Bytes{}
	require.Equal(t, "bytes", o.TypeName())
}

func TestObject_IsFalsy(t *testing.T) {
	var o nanojs.Object = &nanojs.Int{Value: 0}
	require.True(t, o.IsFalsy())
	o = &nanojs.Int{Value: 1}
	require.False(t, o.IsFalsy())
	o = &nanojs.Float{Value: 0}
	require.False(t, o.IsFalsy())
	o = &nanojs.Float{Value: 1}
	require.False(t, o.IsFalsy())
	o = &nanojs.Char{Value: ' '}
	require.False(t, o.IsFalsy())
	o = &nanojs.Char{Value: 'T'}
	require.False(t, o.IsFalsy())
	o = &nanojs.String{Value: ""}
	require.True(t, o.IsFalsy())
	o = &nanojs.String{Value: " "}
	require.False(t, o.IsFalsy())
	o = &nanojs.Array{Value: nil}
	require.True(t, o.IsFalsy())
	o = &nanojs.Array{Value: []nanojs.Object{nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &nanojs.Map{Value: nil}
	require.True(t, o.IsFalsy())
	o = &nanojs.Map{Value: map[string]nanojs.Object{"a": nil}} // nil is not valid but still count as 1 element
	require.False(t, o.IsFalsy())
	o = &nanojs.StringIterator{}
	require.True(t, o.IsFalsy())
	o = &nanojs.ArrayIterator{}
	require.True(t, o.IsFalsy())
	o = &nanojs.MapIterator{}
	require.True(t, o.IsFalsy())
	o = &nanojs.BuiltinFunction{}
	require.False(t, o.IsFalsy())
	o = &nanojs.CompiledFunction{}
	require.False(t, o.IsFalsy())
	o = &nanojs.Undefined{}
	require.True(t, o.IsFalsy())
	o = &nanojs.Error{}
	require.True(t, o.IsFalsy())
	o = &nanojs.Bytes{}
	require.True(t, o.IsFalsy())
	o = &nanojs.Bytes{Value: []byte{1, 2}}
	require.False(t, o.IsFalsy())
}

func TestObject_String(t *testing.T) {
	var o nanojs.Object = &nanojs.Int{Value: 0}
	require.Equal(t, "0", o.String())
	o = &nanojs.Int{Value: 1}
	require.Equal(t, "1", o.String())
	o = &nanojs.Float{Value: 0}
	require.Equal(t, "0", o.String())
	o = &nanojs.Float{Value: 1}
	require.Equal(t, "1", o.String())
	o = &nanojs.Char{Value: ' '}
	require.Equal(t, " ", o.String())
	o = &nanojs.Char{Value: 'T'}
	require.Equal(t, "T", o.String())
	o = &nanojs.String{Value: ""}
	require.Equal(t, `""`, o.String())
	o = &nanojs.String{Value: " "}
	require.Equal(t, `" "`, o.String())
	o = &nanojs.Array{Value: nil}
	require.Equal(t, "[]", o.String())
	o = &nanojs.Map{Value: nil}
	require.Equal(t, "{}", o.String())
	o = &nanojs.Error{Value: nil}
	require.Equal(t, "error", o.String())
	o = &nanojs.Error{Value: &nanojs.String{Value: "error 1"}}
	require.Equal(t, `error: "error 1"`, o.String())
	o = &nanojs.StringIterator{}
	require.Equal(t, "<string-iterator>", o.String())
	o = &nanojs.ArrayIterator{}
	require.Equal(t, "<array-iterator>", o.String())
	o = &nanojs.MapIterator{}
	require.Equal(t, "<map-iterator>", o.String())
	o = &nanojs.Undefined{}
	require.Equal(t, "<undefined>", o.String())
	o = &nanojs.Bytes{}
	require.Equal(t, "", o.String())
	o = &nanojs.Bytes{Value: []byte("foo")}
	require.Equal(t, "foo", o.String())
}

func TestObject_BinaryOp(t *testing.T) {
	var o nanojs.Object = &nanojs.Char{}
	_, err := o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.Bool{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.Map{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.ArrayIterator{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.StringIterator{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.MapIterator{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.BuiltinFunction{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.CompiledFunction{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.Undefined{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
	o = &nanojs.Error{}
	_, err = o.BinaryOp(token.Add, nanojs.UndefinedValue)
	require.Error(t, err)
}

func TestArray_BinaryOp(t *testing.T) {
	testBinaryOp(t, &nanojs.Array{Value: nil}, token.Add,
		&nanojs.Array{Value: nil}, &nanojs.Array{Value: nil})
	testBinaryOp(t, &nanojs.Array{Value: nil}, token.Add,
		&nanojs.Array{Value: []nanojs.Object{}}, &nanojs.Array{Value: nil})
	testBinaryOp(t, &nanojs.Array{Value: []nanojs.Object{}}, token.Add,
		&nanojs.Array{Value: nil}, &nanojs.Array{Value: []nanojs.Object{}})
	testBinaryOp(t, &nanojs.Array{Value: []nanojs.Object{}}, token.Add,
		&nanojs.Array{Value: []nanojs.Object{}},
		&nanojs.Array{Value: []nanojs.Object{}})
	testBinaryOp(t, &nanojs.Array{Value: nil}, token.Add,
		&nanojs.Array{Value: []nanojs.Object{
			&nanojs.Int{Value: 1},
		}}, &nanojs.Array{Value: []nanojs.Object{
			&nanojs.Int{Value: 1},
		}})
	testBinaryOp(t, &nanojs.Array{Value: nil}, token.Add,
		&nanojs.Array{Value: []nanojs.Object{
			&nanojs.Int{Value: 1},
			&nanojs.Int{Value: 2},
			&nanojs.Int{Value: 3},
		}}, &nanojs.Array{Value: []nanojs.Object{
			&nanojs.Int{Value: 1},
			&nanojs.Int{Value: 2},
			&nanojs.Int{Value: 3},
		}})
	testBinaryOp(t, &nanojs.Array{Value: []nanojs.Object{
		&nanojs.Int{Value: 1},
		&nanojs.Int{Value: 2},
		&nanojs.Int{Value: 3},
	}}, token.Add, &nanojs.Array{Value: nil},
		&nanojs.Array{Value: []nanojs.Object{
			&nanojs.Int{Value: 1},
			&nanojs.Int{Value: 2},
			&nanojs.Int{Value: 3},
		}})
	testBinaryOp(t, &nanojs.Array{Value: []nanojs.Object{
		&nanojs.Int{Value: 1},
		&nanojs.Int{Value: 2},
		&nanojs.Int{Value: 3},
	}}, token.Add, &nanojs.Array{Value: []nanojs.Object{
		&nanojs.Int{Value: 4},
		&nanojs.Int{Value: 5},
		&nanojs.Int{Value: 6},
	}}, &nanojs.Array{Value: []nanojs.Object{
		&nanojs.Int{Value: 1},
		&nanojs.Int{Value: 2},
		&nanojs.Int{Value: 3},
		&nanojs.Int{Value: 4},
		&nanojs.Int{Value: 5},
		&nanojs.Int{Value: 6},
	}})
}

func TestError_Equals(t *testing.T) {
	err1 := &nanojs.Error{Value: &nanojs.String{Value: "some error"}}
	err2 := err1
	require.True(t, err1.Equals(err2))
	require.True(t, err2.Equals(err1))

	err2 = &nanojs.Error{Value: &nanojs.String{Value: "some error"}}
	require.False(t, err1.Equals(err2))
	require.False(t, err2.Equals(err1))
}

func TestFloat_BinaryOp(t *testing.T) {
	// float + float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Add,
				&nanojs.Float{Value: r}, &nanojs.Float{Value: l + r})
		}
	}

	// float - float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Sub,
				&nanojs.Float{Value: r}, &nanojs.Float{Value: l - r})
		}
	}

	// float * float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Mul,
				&nanojs.Float{Value: r}, &nanojs.Float{Value: l * r})
		}
	}

	// float / float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			if r != 0 {
				testBinaryOp(t, &nanojs.Float{Value: l}, token.Quo,
					&nanojs.Float{Value: r}, &nanojs.Float{Value: l / r})
			}
		}
	}

	// float < float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Less,
				&nanojs.Float{Value: r}, boolValue(l < r))
		}
	}

	// float > float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Greater,
				&nanojs.Float{Value: r}, boolValue(l > r))
		}
	}

	// float <= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.LessEq,
				&nanojs.Float{Value: r}, boolValue(l <= r))
		}
	}

	// float >= float
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := float64(-2); r <= 2.1; r += 0.4 {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.GreaterEq,
				&nanojs.Float{Value: r}, boolValue(l >= r))
		}
	}

	// float + int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Add,
				&nanojs.Int{Value: r}, &nanojs.Float{Value: l + float64(r)})
		}
	}

	// float - int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Sub,
				&nanojs.Int{Value: r}, &nanojs.Float{Value: l - float64(r)})
		}
	}

	// float * int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Mul,
				&nanojs.Int{Value: r}, &nanojs.Float{Value: l * float64(r)})
		}
	}

	// float / int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &nanojs.Float{Value: l}, token.Quo,
					&nanojs.Int{Value: r},
					&nanojs.Float{Value: l / float64(r)})
			}
		}
	}

	// float < int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Less,
				&nanojs.Int{Value: r}, boolValue(l < float64(r)))
		}
	}

	// float > int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.Greater,
				&nanojs.Int{Value: r}, boolValue(l > float64(r)))
		}
	}

	// float <= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.LessEq,
				&nanojs.Int{Value: r}, boolValue(l <= float64(r)))
		}
	}

	// float >= int
	for l := float64(-2); l <= 2.1; l += 0.4 {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Float{Value: l}, token.GreaterEq,
				&nanojs.Int{Value: r}, boolValue(l >= float64(r)))
		}
	}
}

func TestInt_BinaryOp(t *testing.T) {
	// int + int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Add,
				&nanojs.Int{Value: r}, &nanojs.Int{Value: l + r})
		}
	}

	// int - int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Sub,
				&nanojs.Int{Value: r}, &nanojs.Int{Value: l - r})
		}
	}

	// int * int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Mul,
				&nanojs.Int{Value: r}, &nanojs.Int{Value: l * r})
		}
	}

	// int / int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			if r != 0 {
				testBinaryOp(t, &nanojs.Int{Value: l}, token.Quo,
					&nanojs.Int{Value: r}, &nanojs.Int{Value: l / r})
			}
		}
	}

	// int % int
	for l := int64(-4); l <= 4; l++ {
		for r := -int64(-4); r <= 4; r++ {
			if r == 0 {
				testBinaryOp(t, &nanojs.Int{Value: l}, token.Rem,
					&nanojs.Int{Value: r}, &nanojs.Int{Value: l % r})
			}
		}
	}

	// int & int
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.And, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.And, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(1) & int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.And, &nanojs.Int{Value: 1},
		&nanojs.Int{Value: int64(0) & int64(1)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.And, &nanojs.Int{Value: 1},
		&nanojs.Int{Value: int64(1)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.And, &nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0) & int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.And, &nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1) & int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: int64(0xffffffff)}, token.And,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1984}, token.And,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1984) & int64(0xffffffff)})
	testBinaryOp(t, &nanojs.Int{Value: -1984}, token.And,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(-1984) & int64(0xffffffff)})

	// int | int
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.Or, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.Or, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(1) | int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.Or, &nanojs.Int{Value: 1},
		&nanojs.Int{Value: int64(0) | int64(1)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.Or, &nanojs.Int{Value: 1},
		&nanojs.Int{Value: int64(1)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.Or, &nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0) | int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.Or, &nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1) | int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: int64(0xffffffff)}, token.Or,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1984}, token.Or,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1984) | int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: -1984}, token.Or,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(-1984) | int64(0xffffffff)})

	// int ^ int
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.Xor, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.Xor, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(1) ^ int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.Xor, &nanojs.Int{Value: 1},
		&nanojs.Int{Value: int64(0) ^ int64(1)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.Xor, &nanojs.Int{Value: 1},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.Xor, &nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.Xor, &nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: int64(0xffffffff)}, token.Xor,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1984}, token.Xor,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1984) ^ int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: -1984}, token.Xor,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(-1984) ^ int64(0xffffffff)})

	// int &^ int
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.AndNot, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.AndNot, &nanojs.Int{Value: 0},
		&nanojs.Int{Value: int64(1) &^ int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.AndNot,
		&nanojs.Int{Value: 1}, &nanojs.Int{Value: int64(0) &^ int64(1)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.AndNot, &nanojs.Int{Value: 1},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 0}, token.AndNot,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1}, token.AndNot,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: int64(0xffffffff)}, token.AndNot,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(0)})
	testBinaryOp(t,
		&nanojs.Int{Value: 1984}, token.AndNot,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(1984) &^ int64(0xffffffff)})
	testBinaryOp(t,
		&nanojs.Int{Value: -1984}, token.AndNot,
		&nanojs.Int{Value: int64(0xffffffff)},
		&nanojs.Int{Value: int64(-1984) &^ int64(0xffffffff)})

	// int << int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&nanojs.Int{Value: 0}, token.Shl, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(0) << uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: 1}, token.Shl, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(1) << uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: 2}, token.Shl, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(2) << uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: -1}, token.Shl, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(-1) << uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: -2}, token.Shl, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(-2) << uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: int64(0xffffffff)}, token.Shl,
			&nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(0xffffffff) << uint(s)})
	}

	// int >> int
	for s := int64(0); s < 64; s++ {
		testBinaryOp(t,
			&nanojs.Int{Value: 0}, token.Shr, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(0) >> uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: 1}, token.Shr, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(1) >> uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: 2}, token.Shr, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(2) >> uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: -1}, token.Shr, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(-1) >> uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: -2}, token.Shr, &nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(-2) >> uint(s)})
		testBinaryOp(t,
			&nanojs.Int{Value: int64(0xffffffff)}, token.Shr,
			&nanojs.Int{Value: s},
			&nanojs.Int{Value: int64(0xffffffff) >> uint(s)})
	}

	// int < int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Less,
				&nanojs.Int{Value: r}, boolValue(l < r))
		}
	}

	// int > int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Greater,
				&nanojs.Int{Value: r}, boolValue(l > r))
		}
	}

	// int <= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.LessEq,
				&nanojs.Int{Value: r}, boolValue(l <= r))
		}
	}

	// int >= int
	for l := int64(-2); l <= 2; l++ {
		for r := int64(-2); r <= 2; r++ {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.GreaterEq,
				&nanojs.Int{Value: r}, boolValue(l >= r))
		}
	}

	// int + float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Add,
				&nanojs.Float{Value: r},
				&nanojs.Float{Value: float64(l) + r})
		}
	}

	// int - float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Sub,
				&nanojs.Float{Value: r},
				&nanojs.Float{Value: float64(l) - r})
		}
	}

	// int * float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Mul,
				&nanojs.Float{Value: r},
				&nanojs.Float{Value: float64(l) * r})
		}
	}

	// int / float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			if r != 0 {
				testBinaryOp(t, &nanojs.Int{Value: l}, token.Quo,
					&nanojs.Float{Value: r},
					&nanojs.Float{Value: float64(l) / r})
			}
		}
	}

	// int < float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Less,
				&nanojs.Float{Value: r}, boolValue(float64(l) < r))
		}
	}

	// int > float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.Greater,
				&nanojs.Float{Value: r}, boolValue(float64(l) > r))
		}
	}

	// int <= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.LessEq,
				&nanojs.Float{Value: r}, boolValue(float64(l) <= r))
		}
	}

	// int >= float
	for l := int64(-2); l <= 2; l++ {
		for r := float64(-2); r <= 2.1; r += 0.5 {
			testBinaryOp(t, &nanojs.Int{Value: l}, token.GreaterEq,
				&nanojs.Float{Value: r}, boolValue(float64(l) >= r))
		}
	}
}

func TestMap_Index(t *testing.T) {
	m := &nanojs.Map{Value: make(map[string]nanojs.Object)}
	k := &nanojs.Int{Value: 1}
	v := &nanojs.String{Value: "abcdef"}
	err := m.IndexSet(k, v)

	require.NoError(t, err)

	res, err := m.IndexGet(k)
	require.NoError(t, err)
	require.Equal(t, v, res)
}

func TestString_BinaryOp(t *testing.T) {
	lstr := "abcde"
	rstr := "01234"
	for l := 0; l < len(lstr); l++ {
		for r := 0; r < len(rstr); r++ {
			ls := lstr[l:]
			rs := rstr[r:]
			testBinaryOp(t, &nanojs.String{Value: ls}, token.Add,
				&nanojs.String{Value: rs},
				&nanojs.String{Value: ls + rs})

			rc := []rune(rstr)[r]
			testBinaryOp(t, &nanojs.String{Value: ls}, token.Add,
				&nanojs.Char{Value: rc},
				&nanojs.String{Value: ls + string(rc)})
		}
	}
}

func testBinaryOp(
	t *testing.T,
	lhs nanojs.Object,
	op token.Token,
	rhs nanojs.Object,
	expected nanojs.Object,
) {
	t.Helper()
	actual, err := lhs.BinaryOp(op, rhs)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func boolValue(b bool) nanojs.Object {
	if b {
		return nanojs.TrueValue
	}
	return nanojs.FalseValue
}
