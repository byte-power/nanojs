package stdlib

import (
	"github.com/zeaphoo/nanojs/v2"
)

func wrapError(err error) nanojs.Object {
	if err == nil {
		return nanojs.TrueValue
	}
	return &nanojs.Error{Value: &nanojs.String{Value: err.Error()}}
}
