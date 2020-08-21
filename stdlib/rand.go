package stdlib

import (
	"math/rand"

	"github.com/zeaphoo/nanojs/v2"
)

var randModule = map[string]nanojs.Object{
	"int": &nanojs.UserFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &nanojs.UserFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &nanojs.UserFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &nanojs.UserFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &nanojs.UserFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &nanojs.UserFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &nanojs.UserFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &nanojs.UserFunction{
		Name: "read",
		Value: func(args ...nanojs.Object) (ret nanojs.Object, err error) {
			if len(args) != 1 {
				return nil, nanojs.ErrWrongNumArguments
			}
			y1, ok := args[0].(*nanojs.Bytes)
			if !ok {
				return nil, nanojs.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "bytes",
					Found:    args[0].TypeName(),
				}
			}
			res, err := rand.Read(y1.Value)
			if err != nil {
				ret = wrapError(err)
				return
			}
			return &nanojs.Int{Value: int64(res)}, nil
		},
	},
	"rand": &nanojs.UserFunction{
		Name: "rand",
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
			src := rand.NewSource(i1)
			return randRand(rand.New(src)), nil
		},
	},
}

func randRand(r *rand.Rand) *nanojs.ImmutableMap {
	return &nanojs.ImmutableMap{
		Value: map[string]nanojs.Object{
			"int": &nanojs.UserFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &nanojs.UserFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &nanojs.UserFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &nanojs.UserFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &nanojs.UserFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &nanojs.UserFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &nanojs.UserFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &nanojs.UserFunction{
				Name: "read",
				Value: func(args ...nanojs.Object) (
					ret nanojs.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, nanojs.ErrWrongNumArguments
					}
					y1, ok := args[0].(*nanojs.Bytes)
					if !ok {
						return nil, nanojs.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &nanojs.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
