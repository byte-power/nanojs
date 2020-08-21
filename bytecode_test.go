package nanojs_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/zeaphoo/nanojs/v2"
	"github.com/zeaphoo/nanojs/v2/parser"
	"github.com/zeaphoo/nanojs/v2/require"
)

type srcfile struct {
	name string
	size int
}

func TestBytecode(t *testing.T) {
	testBytecodeSerialization(t, bytecode(concatInsts(), objectsArray()))

	testBytecodeSerialization(t, bytecode(
		concatInsts(), objectsArray(
			&nanojs.Char{Value: 'y'},
			&nanojs.Float{Value: 93.11},
			compiledFunction(1, 0,
				nanojs.MakeInstruction(parser.OpConstant, 3),
				nanojs.MakeInstruction(parser.OpSetLocal, 0),
				nanojs.MakeInstruction(parser.OpGetGlobal, 0),
				nanojs.MakeInstruction(parser.OpGetFree, 0)),
			&nanojs.Float{Value: 39.2},
			&nanojs.Int{Value: 192},
			&nanojs.String{Value: "bar"})))

	testBytecodeSerialization(t, bytecodeFileSet(
		concatInsts(
			nanojs.MakeInstruction(parser.OpConstant, 0),
			nanojs.MakeInstruction(parser.OpSetGlobal, 0),
			nanojs.MakeInstruction(parser.OpConstant, 6),
			nanojs.MakeInstruction(parser.OpPop)),
		objectsArray(
			&nanojs.Int{Value: 55},
			&nanojs.Int{Value: 66},
			&nanojs.Int{Value: 77},
			&nanojs.Int{Value: 88},
			&nanojs.ImmutableMap{
				Value: map[string]nanojs.Object{
					"array": &nanojs.ImmutableArray{
						Value: []nanojs.Object{
							&nanojs.Int{Value: 1},
							&nanojs.Int{Value: 2},
							&nanojs.Int{Value: 3},
							nanojs.TrueValue,
							nanojs.FalseValue,
							nanojs.UndefinedValue,
						},
					},
					"true":  nanojs.TrueValue,
					"false": nanojs.FalseValue,
					"bytes": &nanojs.Bytes{Value: make([]byte, 16)},
					"char":  &nanojs.Char{Value: 'Y'},
					"error": &nanojs.Error{Value: &nanojs.String{
						Value: "some error",
					}},
					"float": &nanojs.Float{Value: -19.84},
					"immutable_array": &nanojs.ImmutableArray{
						Value: []nanojs.Object{
							&nanojs.Int{Value: 1},
							&nanojs.Int{Value: 2},
							&nanojs.Int{Value: 3},
							nanojs.TrueValue,
							nanojs.FalseValue,
							nanojs.UndefinedValue,
						},
					},
					"immutable_map": &nanojs.ImmutableMap{
						Value: map[string]nanojs.Object{
							"a": &nanojs.Int{Value: 1},
							"b": &nanojs.Int{Value: 2},
							"c": &nanojs.Int{Value: 3},
							"d": nanojs.TrueValue,
							"e": nanojs.FalseValue,
							"f": nanojs.UndefinedValue,
						},
					},
					"int": &nanojs.Int{Value: 91},
					"map": &nanojs.Map{
						Value: map[string]nanojs.Object{
							"a": &nanojs.Int{Value: 1},
							"b": &nanojs.Int{Value: 2},
							"c": &nanojs.Int{Value: 3},
							"d": nanojs.TrueValue,
							"e": nanojs.FalseValue,
							"f": nanojs.UndefinedValue,
						},
					},
					"string":    &nanojs.String{Value: "foo bar"},
					"time":      &nanojs.Time{Value: time.Now()},
					"undefined": nanojs.UndefinedValue,
				},
			},
			compiledFunction(1, 0,
				nanojs.MakeInstruction(parser.OpConstant, 3),
				nanojs.MakeInstruction(parser.OpSetLocal, 0),
				nanojs.MakeInstruction(parser.OpGetGlobal, 0),
				nanojs.MakeInstruction(parser.OpGetFree, 0),
				nanojs.MakeInstruction(parser.OpBinaryOp, 11),
				nanojs.MakeInstruction(parser.OpGetFree, 1),
				nanojs.MakeInstruction(parser.OpBinaryOp, 11),
				nanojs.MakeInstruction(parser.OpGetLocal, 0),
				nanojs.MakeInstruction(parser.OpBinaryOp, 11),
				nanojs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				nanojs.MakeInstruction(parser.OpConstant, 2),
				nanojs.MakeInstruction(parser.OpSetLocal, 0),
				nanojs.MakeInstruction(parser.OpGetFree, 0),
				nanojs.MakeInstruction(parser.OpGetLocal, 0),
				nanojs.MakeInstruction(parser.OpClosure, 4, 2),
				nanojs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				nanojs.MakeInstruction(parser.OpConstant, 1),
				nanojs.MakeInstruction(parser.OpSetLocal, 0),
				nanojs.MakeInstruction(parser.OpGetLocal, 0),
				nanojs.MakeInstruction(parser.OpClosure, 5, 1),
				nanojs.MakeInstruction(parser.OpReturn, 1))),
		fileSet(srcfile{name: "file1", size: 100},
			srcfile{name: "file2", size: 200})))
}

func TestBytecode_RemoveDuplicates(t *testing.T) {
	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(), objectsArray(
				&nanojs.Char{Value: 'y'},
				&nanojs.Float{Value: 93.11},
				compiledFunction(1, 0,
					nanojs.MakeInstruction(parser.OpConstant, 3),
					nanojs.MakeInstruction(parser.OpSetLocal, 0),
					nanojs.MakeInstruction(parser.OpGetGlobal, 0),
					nanojs.MakeInstruction(parser.OpGetFree, 0)),
				&nanojs.Float{Value: 39.2},
				&nanojs.Int{Value: 192},
				&nanojs.String{Value: "bar"})),
		bytecode(
			concatInsts(), objectsArray(
				&nanojs.Char{Value: 'y'},
				&nanojs.Float{Value: 93.11},
				compiledFunction(1, 0,
					nanojs.MakeInstruction(parser.OpConstant, 3),
					nanojs.MakeInstruction(parser.OpSetLocal, 0),
					nanojs.MakeInstruction(parser.OpGetGlobal, 0),
					nanojs.MakeInstruction(parser.OpGetFree, 0)),
				&nanojs.Float{Value: 39.2},
				&nanojs.Int{Value: 192},
				&nanojs.String{Value: "bar"})))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				nanojs.MakeInstruction(parser.OpConstant, 0),
				nanojs.MakeInstruction(parser.OpConstant, 1),
				nanojs.MakeInstruction(parser.OpConstant, 2),
				nanojs.MakeInstruction(parser.OpConstant, 3),
				nanojs.MakeInstruction(parser.OpConstant, 4),
				nanojs.MakeInstruction(parser.OpConstant, 5),
				nanojs.MakeInstruction(parser.OpConstant, 6),
				nanojs.MakeInstruction(parser.OpConstant, 7),
				nanojs.MakeInstruction(parser.OpConstant, 8),
				nanojs.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&nanojs.Int{Value: 1},
				&nanojs.Float{Value: 2.0},
				&nanojs.Char{Value: '3'},
				&nanojs.String{Value: "four"},
				compiledFunction(1, 0,
					nanojs.MakeInstruction(parser.OpConstant, 3),
					nanojs.MakeInstruction(parser.OpConstant, 7),
					nanojs.MakeInstruction(parser.OpSetLocal, 0),
					nanojs.MakeInstruction(parser.OpGetGlobal, 0),
					nanojs.MakeInstruction(parser.OpGetFree, 0)),
				&nanojs.Int{Value: 1},
				&nanojs.Float{Value: 2.0},
				&nanojs.Char{Value: '3'},
				&nanojs.String{Value: "four"})),
		bytecode(
			concatInsts(
				nanojs.MakeInstruction(parser.OpConstant, 0),
				nanojs.MakeInstruction(parser.OpConstant, 1),
				nanojs.MakeInstruction(parser.OpConstant, 2),
				nanojs.MakeInstruction(parser.OpConstant, 3),
				nanojs.MakeInstruction(parser.OpConstant, 4),
				nanojs.MakeInstruction(parser.OpConstant, 0),
				nanojs.MakeInstruction(parser.OpConstant, 1),
				nanojs.MakeInstruction(parser.OpConstant, 2),
				nanojs.MakeInstruction(parser.OpConstant, 3),
				nanojs.MakeInstruction(parser.OpClosure, 4, 1)),
			objectsArray(
				&nanojs.Int{Value: 1},
				&nanojs.Float{Value: 2.0},
				&nanojs.Char{Value: '3'},
				&nanojs.String{Value: "four"},
				compiledFunction(1, 0,
					nanojs.MakeInstruction(parser.OpConstant, 3),
					nanojs.MakeInstruction(parser.OpConstant, 2),
					nanojs.MakeInstruction(parser.OpSetLocal, 0),
					nanojs.MakeInstruction(parser.OpGetGlobal, 0),
					nanojs.MakeInstruction(parser.OpGetFree, 0)))))

	testBytecodeRemoveDuplicates(t,
		bytecode(
			concatInsts(
				nanojs.MakeInstruction(parser.OpConstant, 0),
				nanojs.MakeInstruction(parser.OpConstant, 1),
				nanojs.MakeInstruction(parser.OpConstant, 2),
				nanojs.MakeInstruction(parser.OpConstant, 3),
				nanojs.MakeInstruction(parser.OpConstant, 4)),
			objectsArray(
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2},
				&nanojs.Int{Value: 3},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 3})),
		bytecode(
			concatInsts(
				nanojs.MakeInstruction(parser.OpConstant, 0),
				nanojs.MakeInstruction(parser.OpConstant, 1),
				nanojs.MakeInstruction(parser.OpConstant, 2),
				nanojs.MakeInstruction(parser.OpConstant, 0),
				nanojs.MakeInstruction(parser.OpConstant, 2)),
			objectsArray(
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2},
				&nanojs.Int{Value: 3})))
}

func TestBytecode_CountObjects(t *testing.T) {
	b := bytecode(
		concatInsts(),
		objectsArray(
			&nanojs.Int{Value: 55},
			&nanojs.Int{Value: 66},
			&nanojs.Int{Value: 77},
			&nanojs.Int{Value: 88},
			compiledFunction(1, 0,
				nanojs.MakeInstruction(parser.OpConstant, 3),
				nanojs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				nanojs.MakeInstruction(parser.OpConstant, 2),
				nanojs.MakeInstruction(parser.OpReturn, 1)),
			compiledFunction(1, 0,
				nanojs.MakeInstruction(parser.OpConstant, 1),
				nanojs.MakeInstruction(parser.OpReturn, 1))))
	require.Equal(t, 7, b.CountObjects())
}

func fileSet(files ...srcfile) *parser.SourceFileSet {
	fileSet := parser.NewFileSet()
	for _, f := range files {
		fileSet.AddFile(f.name, -1, f.size)
	}
	return fileSet
}

func bytecodeFileSet(
	instructions []byte,
	constants []nanojs.Object,
	fileSet *parser.SourceFileSet,
) *nanojs.Bytecode {
	return &nanojs.Bytecode{
		FileSet:      fileSet,
		MainFunction: &nanojs.CompiledFunction{Instructions: instructions},
		Constants:    constants,
	}
}

func testBytecodeRemoveDuplicates(
	t *testing.T,
	input, expected *nanojs.Bytecode,
) {
	input.RemoveDuplicates()

	require.Equal(t, expected.FileSet, input.FileSet)
	require.Equal(t, expected.MainFunction, input.MainFunction)
	require.Equal(t, expected.Constants, input.Constants)
}

func testBytecodeSerialization(t *testing.T, b *nanojs.Bytecode) {
	var buf bytes.Buffer
	err := b.Encode(&buf)
	require.NoError(t, err)

	r := &nanojs.Bytecode{}
	err = r.Decode(bytes.NewReader(buf.Bytes()), nil)
	require.NoError(t, err)

	require.Equal(t, b.FileSet, r.FileSet)
	require.Equal(t, b.MainFunction, r.MainFunction)
	require.Equal(t, b.Constants, r.Constants)
}
