package stdlib

import (
	"encoding/hex"

	"github.com/zeaphoo/nanojs/v2"
)

var hexModule = map[string]nanojs.Object{
	"encode": &nanojs.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &nanojs.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
