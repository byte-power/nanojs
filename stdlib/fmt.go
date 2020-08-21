package stdlib

import (
	"fmt"

	"github.com/zeaphoo/nanojs/v2"
)

var fmtModule = map[string]nanojs.Object{
	"print":   &nanojs.UserFunction{Name: "print", Value: fmtPrint},
	"printf":  &nanojs.UserFunction{Name: "printf", Value: fmtPrintf},
	"println": &nanojs.UserFunction{Name: "println", Value: fmtPrintln},
	"sprintf": &nanojs.UserFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(args ...nanojs.Object) (ret nanojs.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtPrintf(args ...nanojs.Object) (ret nanojs.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, nanojs.ErrWrongNumArguments
	}

	format, ok := args[0].(*nanojs.String)
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Print(format)
		return nil, nil
	}

	s, err := nanojs.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Print(s)
	return nil, nil
}

func fmtPrintln(args ...nanojs.Object) (ret nanojs.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtSprintf(args ...nanojs.Object) (ret nanojs.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, nanojs.ErrWrongNumArguments
	}

	format, ok := args[0].(*nanojs.String)
	if !ok {
		return nil, nanojs.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := nanojs.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &nanojs.String{Value: s}, nil
}

func getPrintArgs(args ...nanojs.Object) ([]interface{}, error) {
	var printArgs []interface{}
	l := 0
	for _, arg := range args {
		s, _ := nanojs.ToString(arg)
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > nanojs.MaxStringLen {
			return nil, nanojs.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
