package stdlib

import (
	"os/exec"

	"github.com/zeaphoo/nanojs/v2"
)

func makeOSExecCommand(cmd *exec.Cmd) *nanojs.ImmutableMap {
	return &nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			// combined_output() => bytes/error
			"combined_output": &nanojs.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &nanojs.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &nanojs.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &nanojs.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &nanojs.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &nanojs.UserFunction{
				Name: "set_path",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 1 {
						return nil, nanojs.ErrWrongNumArguments
					}
					s1, ok := nanojs.ToString(args[0])
					if !ok {
						return nil, nanojs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Path = s1
					return nanojs.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &nanojs.UserFunction{
				Name: "set_dir",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 1 {
						return nil, nanojs.ErrWrongNumArguments
					}
					s1, ok := nanojs.ToString(args[0])
					if !ok {
						return nil, nanojs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Dir = s1
					return nanojs.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &nanojs.UserFunction{
				Name: "set_env",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 1 {
						return nil, nanojs.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *nanojs.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *nanojs.ImmutableArray:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, nanojs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return nanojs.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &nanojs.UserFunction{
				Name: "process",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 0 {
						return nil, nanojs.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
