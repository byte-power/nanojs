package stdlib

import (
	"math"

	"github.com/zeaphoo/nanojs/v2"
)

var mathModule = map[string]nanojs.Object{
	"e":       &nanojs.Float{Value: math.E},
	"pi":      &nanojs.Float{Value: math.Pi},
	"phi":     &nanojs.Float{Value: math.Phi},
	"sqrt2":   &nanojs.Float{Value: math.Sqrt2},
	"sqrtE":   &nanojs.Float{Value: math.SqrtE},
	"sqrtPi":  &nanojs.Float{Value: math.SqrtPi},
	"sqrtPhi": &nanojs.Float{Value: math.SqrtPhi},
	"ln2":     &nanojs.Float{Value: math.Ln2},
	"log2E":   &nanojs.Float{Value: math.Log2E},
	"ln10":    &nanojs.Float{Value: math.Ln10},
	"log10E":  &nanojs.Float{Value: math.Log10E},
	"abs": &nanojs.UserFunction{
		Name:  "abs",
		Value: FuncAFRF(math.Abs),
	},
	"acos": &nanojs.UserFunction{
		Name:  "acos",
		Value: FuncAFRF(math.Acos),
	},
	"acosh": &nanojs.UserFunction{
		Name:  "acosh",
		Value: FuncAFRF(math.Acosh),
	},
	"asin": &nanojs.UserFunction{
		Name:  "asin",
		Value: FuncAFRF(math.Asin),
	},
	"asinh": &nanojs.UserFunction{
		Name:  "asinh",
		Value: FuncAFRF(math.Asinh),
	},
	"atan": &nanojs.UserFunction{
		Name:  "atan",
		Value: FuncAFRF(math.Atan),
	},
	"atan2": &nanojs.UserFunction{
		Name:  "atan2",
		Value: FuncAFFRF(math.Atan2),
	},
	"atanh": &nanojs.UserFunction{
		Name:  "atanh",
		Value: FuncAFRF(math.Atanh),
	},
	"cbrt": &nanojs.UserFunction{
		Name:  "cbrt",
		Value: FuncAFRF(math.Cbrt),
	},
	"ceil": &nanojs.UserFunction{
		Name:  "ceil",
		Value: FuncAFRF(math.Ceil),
	},
	"copysign": &nanojs.UserFunction{
		Name:  "copysign",
		Value: FuncAFFRF(math.Copysign),
	},
	"cos": &nanojs.UserFunction{
		Name:  "cos",
		Value: FuncAFRF(math.Cos),
	},
	"cosh": &nanojs.UserFunction{
		Name:  "cosh",
		Value: FuncAFRF(math.Cosh),
	},
	"dim": &nanojs.UserFunction{
		Name:  "dim",
		Value: FuncAFFRF(math.Dim),
	},
	"erf": &nanojs.UserFunction{
		Name:  "erf",
		Value: FuncAFRF(math.Erf),
	},
	"erfc": &nanojs.UserFunction{
		Name:  "erfc",
		Value: FuncAFRF(math.Erfc),
	},
	"exp": &nanojs.UserFunction{
		Name:  "exp",
		Value: FuncAFRF(math.Exp),
	},
	"exp2": &nanojs.UserFunction{
		Name:  "exp2",
		Value: FuncAFRF(math.Exp2),
	},
	"expm1": &nanojs.UserFunction{
		Name:  "expm1",
		Value: FuncAFRF(math.Expm1),
	},
	"floor": &nanojs.UserFunction{
		Name:  "floor",
		Value: FuncAFRF(math.Floor),
	},
	"gamma": &nanojs.UserFunction{
		Name:  "gamma",
		Value: FuncAFRF(math.Gamma),
	},
	"hypot": &nanojs.UserFunction{
		Name:  "hypot",
		Value: FuncAFFRF(math.Hypot),
	},
	"ilogb": &nanojs.UserFunction{
		Name:  "ilogb",
		Value: FuncAFRI(math.Ilogb),
	},
	"inf": &nanojs.UserFunction{
		Name:  "inf",
		Value: FuncAIRF(math.Inf),
	},
	"is_inf": &nanojs.UserFunction{
		Name:  "is_inf",
		Value: FuncAFIRB(math.IsInf),
	},
	"is_nan": &nanojs.UserFunction{
		Name:  "is_nan",
		Value: FuncAFRB(math.IsNaN),
	},
	"j0": &nanojs.UserFunction{
		Name:  "j0",
		Value: FuncAFRF(math.J0),
	},
	"j1": &nanojs.UserFunction{
		Name:  "j1",
		Value: FuncAFRF(math.J1),
	},
	"jn": &nanojs.UserFunction{
		Name:  "jn",
		Value: FuncAIFRF(math.Jn),
	},
	"ldexp": &nanojs.UserFunction{
		Name:  "ldexp",
		Value: FuncAFIRF(math.Ldexp),
	},
	"log": &nanojs.UserFunction{
		Name:  "log",
		Value: FuncAFRF(math.Log),
	},
	"log10": &nanojs.UserFunction{
		Name:  "log10",
		Value: FuncAFRF(math.Log10),
	},
	"log1p": &nanojs.UserFunction{
		Name:  "log1p",
		Value: FuncAFRF(math.Log1p),
	},
	"log2": &nanojs.UserFunction{
		Name:  "log2",
		Value: FuncAFRF(math.Log2),
	},
	"logb": &nanojs.UserFunction{
		Name:  "logb",
		Value: FuncAFRF(math.Logb),
	},
	"max": &nanojs.UserFunction{
		Name:  "max",
		Value: FuncAFFRF(math.Max),
	},
	"min": &nanojs.UserFunction{
		Name:  "min",
		Value: FuncAFFRF(math.Min),
	},
	"mod": &nanojs.UserFunction{
		Name:  "mod",
		Value: FuncAFFRF(math.Mod),
	},
	"nan": &nanojs.UserFunction{
		Name:  "nan",
		Value: FuncARF(math.NaN),
	},
	"nextafter": &nanojs.UserFunction{
		Name:  "nextafter",
		Value: FuncAFFRF(math.Nextafter),
	},
	"pow": &nanojs.UserFunction{
		Name:  "pow",
		Value: FuncAFFRF(math.Pow),
	},
	"pow10": &nanojs.UserFunction{
		Name:  "pow10",
		Value: FuncAIRF(math.Pow10),
	},
	"remainder": &nanojs.UserFunction{
		Name:  "remainder",
		Value: FuncAFFRF(math.Remainder),
	},
	"signbit": &nanojs.UserFunction{
		Name:  "signbit",
		Value: FuncAFRB(math.Signbit),
	},
	"sin": &nanojs.UserFunction{
		Name:  "sin",
		Value: FuncAFRF(math.Sin),
	},
	"sinh": &nanojs.UserFunction{
		Name:  "sinh",
		Value: FuncAFRF(math.Sinh),
	},
	"sqrt": &nanojs.UserFunction{
		Name:  "sqrt",
		Value: FuncAFRF(math.Sqrt),
	},
	"tan": &nanojs.UserFunction{
		Name:  "tan",
		Value: FuncAFRF(math.Tan),
	},
	"tanh": &nanojs.UserFunction{
		Name:  "tanh",
		Value: FuncAFRF(math.Tanh),
	},
	"trunc": &nanojs.UserFunction{
		Name:  "trunc",
		Value: FuncAFRF(math.Trunc),
	},
	"y0": &nanojs.UserFunction{
		Name:  "y0",
		Value: FuncAFRF(math.Y0),
	},
	"y1": &nanojs.UserFunction{
		Name:  "y1",
		Value: FuncAFRF(math.Y1),
	},
	"yn": &nanojs.UserFunction{
		Name:  "yn",
		Value: FuncAIFRF(math.Yn),
	},
}
