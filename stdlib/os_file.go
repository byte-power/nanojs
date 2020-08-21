package stdlib

import (
	"os"

	"github.com/zeaphoo/nanojs/v2"
)

func makeOSFile(file *os.File) *nanojs.ImmutableMap {
	return &nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			// chdir() => true/error
			"chdir": &nanojs.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &nanojs.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &nanojs.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &nanojs.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &nanojs.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &nanojs.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &nanojs.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &nanojs.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &nanojs.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &nanojs.UserFunction{
				Name: "chmod",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 1 {
						return nil, nanojs.ErrWrongNumArguments
					}
					i1, ok := nanojs.ToInt64(args[0])
					if !ok {
						return nil, nanojs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &nanojs.UserFunction{
				Name: "seek",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 2 {
						return nil, nanojs.ErrWrongNumArguments
					}
					i1, ok := nanojs.ToInt64(args[0])
					if !ok {
						return nil, nanojs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					i2, ok := nanojs.ToInt(args[1])
					if !ok {
						return nil, nanojs.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &nanojs.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &nanojs.UserFunction{
				Name: "stat",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 0 {
						return nil, nanojs.ErrWrongNumArguments
					}
					return osStat(&nanojs.String{Value: file.Name()})
				},
			},
		},
	}
}
