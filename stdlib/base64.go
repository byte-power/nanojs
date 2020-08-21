package stdlib

import (
	"encoding/base64"

	"github.com/zeaphoo/nanojs/v2"
)

var base64Module = map[string]nanojs.Object{
	"encode": &nanojs.UserFunction{
		Value: FuncAYRS(base64.StdEncoding.EncodeToString),
	},
	"decode": &nanojs.UserFunction{
		Value: FuncASRYE(base64.StdEncoding.DecodeString),
	},
	"raw_encode": &nanojs.UserFunction{
		Value: FuncAYRS(base64.RawStdEncoding.EncodeToString),
	},
	"raw_decode": &nanojs.UserFunction{
		Value: FuncASRYE(base64.RawStdEncoding.DecodeString),
	},
	"url_encode": &nanojs.UserFunction{
		Value: FuncAYRS(base64.URLEncoding.EncodeToString),
	},
	"url_decode": &nanojs.UserFunction{
		Value: FuncASRYE(base64.URLEncoding.DecodeString),
	},
	"raw_url_encode": &nanojs.UserFunction{
		Value: FuncAYRS(base64.RawURLEncoding.EncodeToString),
	},
	"raw_url_decode": &nanojs.UserFunction{
		Value: FuncASRYE(base64.RawURLEncoding.DecodeString),
	},
}
