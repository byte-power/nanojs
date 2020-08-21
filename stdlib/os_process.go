package stdlib

import (
	"os"
	"syscall"

	"github.com/zeaphoo/nanojs/v2"
)

func makeOSProcessState(state *os.ProcessState) *nanojs.ImmutableMap {
	return &nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			"exited": &nanojs.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &nanojs.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &nanojs.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &nanojs.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *nanojs.ImmutableMap {
	return &nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			"kill": &nanojs.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &nanojs.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &nanojs.UserFunction{
				Name: "signal",
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
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &nanojs.UserFunction{
				Name: "wait",
				Value: func(args ...nanojs.Object) (nanojs.Object, error) {
					if len(args) != 0 {
						return nil, nanojs.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
