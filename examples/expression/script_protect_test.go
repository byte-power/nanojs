package main

import (
	"testing"

	"github.com/zeaphoo/nanojs/v2"
)

const expr = "a.a=fn(a.a)"
const originAA = int64(9)
var script = nanojs.NewScript([]byte(expr))

func init() {
	script.Add("a", map[string]interface{}{"a": originAA})
	script.Add("fn", func(args ...nanojs.Object) (ret nanojs.Object, err error) {
		if len(args) > 0 {
			switch arg := args[0].(type) {
			case *nanojs.Int:
				return &nanojs.Int{Value: arg.Value * 2}, nil
			}
		}
		return &nanojs.Int{Value: 0}, nil
	})
}

func TestRun(t *testing.T) {
	compiled, err := script.Compile()
	if err != nil {
		t.Fatal(err)
	}
	err = compiled.Run()
	if err != nil {
		t.Fatal(err)
	}
	a := compiled.Get("a")
	if v := a.Map()["a"]; v == originAA {
		t.Fatal("a.a should changed but not", v)
	}
}

func BenchmarkRun(b *testing.B) {
	compiled, err := script.Compile()
	if err != nil {
		panic(err)
	}
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			err := compiled.Run()
			if err != nil {
				panic(err)
			}
		}
	})
}

func TestRunOnProtectMode(t *testing.T) {
	compiled, err := script.Compile()
	if err != nil {
		t.Fatal(err)
	}
	err = compiled.RunOnProtectMode()
	if err != nil {
		t.Fatal(err)
	}
	a := compiled.Get("a")
	if v := a.Map()["a"]; v != originAA {
		t.Fatal("a.a should not changed", v)
	}
}

func BenchmarkRunOnProtectMode(b *testing.B) {
	compiled, err := script.Compile()
	if err != nil {
		panic(err)
	}
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			err := compiled.RunOnProtectModeWithCachedVM()
			if err != nil {
				panic(err)
			}
		}
	})
}
