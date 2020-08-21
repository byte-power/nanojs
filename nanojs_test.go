package nanojs_test

import (
	"strings"
	"testing"
	"time"

	"github.com/zeaphoo/nanojs/v2"
	"github.com/zeaphoo/nanojs/v2/parser"
	"github.com/zeaphoo/nanojs/v2/require"
)

func TestInstructions_String(t *testing.T) {
	assertInstructionString(t,
		[][]byte{
			nanojs.MakeInstruction(parser.OpConstant, 1),
			nanojs.MakeInstruction(parser.OpConstant, 2),
			nanojs.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 CONST   1
0003 CONST   2
0006 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			nanojs.MakeInstruction(parser.OpBinaryOp, 11),
			nanojs.MakeInstruction(parser.OpConstant, 2),
			nanojs.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11
0002 CONST   2
0005 CONST   65535`)

	assertInstructionString(t,
		[][]byte{
			nanojs.MakeInstruction(parser.OpBinaryOp, 11),
			nanojs.MakeInstruction(parser.OpGetLocal, 1),
			nanojs.MakeInstruction(parser.OpConstant, 2),
			nanojs.MakeInstruction(parser.OpConstant, 65535),
		},
		`0000 BINARYOP 11
0002 GETL    1
0004 CONST   2
0007 CONST   65535`)
}

func TestMakeInstruction(t *testing.T) {
	makeInstruction(t, []byte{parser.OpConstant, 0, 0},
		parser.OpConstant, 0)
	makeInstruction(t, []byte{parser.OpConstant, 0, 1},
		parser.OpConstant, 1)
	makeInstruction(t, []byte{parser.OpConstant, 255, 254},
		parser.OpConstant, 65534)
	makeInstruction(t, []byte{parser.OpPop}, parser.OpPop)
	makeInstruction(t, []byte{parser.OpTrue}, parser.OpTrue)
	makeInstruction(t, []byte{parser.OpFalse}, parser.OpFalse)
}

func TestNumObjects(t *testing.T) {
	testCountObjects(t, &nanojs.Array{}, 1)
	testCountObjects(t, &nanojs.Array{Value: []nanojs.Object{
		&nanojs.Int{Value: 1},
		&nanojs.Int{Value: 2},
		&nanojs.Array{Value: []nanojs.Object{
			&nanojs.Int{Value: 3},
			&nanojs.Int{Value: 4},
			&nanojs.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, nanojs.TrueValue, 1)
	testCountObjects(t, nanojs.FalseValue, 1)
	testCountObjects(t, &nanojs.BuiltinFunction{}, 1)
	testCountObjects(t, &nanojs.Bytes{Value: []byte("foobar")}, 1)
	testCountObjects(t, &nanojs.Char{Value: '가'}, 1)
	testCountObjects(t, &nanojs.CompiledFunction{}, 1)
	testCountObjects(t, &nanojs.Error{Value: &nanojs.Int{Value: 5}}, 2)
	testCountObjects(t, &nanojs.Float{Value: 19.84}, 1)
	testCountObjects(t, &nanojs.ImmutableArray{Value: []nanojs.Object{
		&nanojs.Int{Value: 1},
		&nanojs.Int{Value: 2},
		&nanojs.ImmutableArray{Value: []nanojs.Object{
			&nanojs.Int{Value: 3},
			&nanojs.Int{Value: 4},
			&nanojs.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			"k1": &nanojs.Int{Value: 1},
			"k2": &nanojs.Int{Value: 2},
			"k3": &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 3},
				&nanojs.Int{Value: 4},
				&nanojs.Int{Value: 5},
			}},
		}}, 7)
	testCountObjects(t, &nanojs.Int{Value: 1984}, 1)
	testCountObjects(t, &nanojs.Map{Value: map[string]nanojs.Object{
		"k1": &nanojs.Int{Value: 1},
		"k2": &nanojs.Int{Value: 2},
		"k3": &nanojs.Array{Value: []nanojs.Object{
			&nanojs.Int{Value: 3},
			&nanojs.Int{Value: 4},
			&nanojs.Int{Value: 5},
		}},
	}}, 7)
	testCountObjects(t, &nanojs.String{Value: "foo bar"}, 1)
	testCountObjects(t, &nanojs.Time{Value: time.Now()}, 1)
	testCountObjects(t, nanojs.UndefinedValue, 1)
}

func testCountObjects(t *testing.T, o nanojs.Object, expected int) {
	require.Equal(t, expected, nanojs.CountObjects(o))
}

func assertInstructionString(
	t *testing.T,
	instructions [][]byte,
	expected string,
) {
	concatted := make([]byte, 0)
	for _, e := range instructions {
		concatted = append(concatted, e...)
	}
	require.Equal(t, expected, strings.Join(
		nanojs.FormatInstructions(concatted, 0), "\n"))
}

func makeInstruction(
	t *testing.T,
	expected []byte,
	opcode parser.Opcode,
	operands ...int,
) {
	inst := nanojs.MakeInstruction(opcode, operands...)
	require.Equal(t, expected, inst)
}
