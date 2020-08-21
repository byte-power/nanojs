package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"github.com/zeaphoo/nanojs/v2"
	"github.com/zeaphoo/nanojs/v2/stdlib/json"
)

var jsonModule = map[string]nanojs.Object{
	"decode": &nanojs.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &nanojs.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &nanojs.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &nanojs.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *nanojs.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &nanojs.Error{
				Value: &nanojs.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *nanojs.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &nanojs.Error{
				Value: &nanojs.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &nanojs.Error{Value: &nanojs.String{Value: err.Error()}}, nil
	}

	return &nanojs.Bytes{Value: b}, nil
}

func jsonIndent(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 3 {
		return nil, nanojs.ErrWrongNumArguments
	}

	prefix, ok := nanojs.ToString(args[1])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := nanojs.ToString(args[2])
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *nanojs.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &nanojs.Error{
				Value: &nanojs.String{Value: err.Error()},
			}, nil
		}
		return &nanojs.Bytes{Value: dst.Bytes()}, nil
	case *nanojs.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &nanojs.Error{
				Value: &nanojs.String{Value: err.Error()},
			}, nil
		}
		return &nanojs.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		return nil, nanojs.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *nanojs.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &nanojs.Bytes{Value: dst.Bytes()}, nil
	case *nanojs.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &nanojs.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
