package nanojs_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/zeaphoo/nanojs/v2"
)

func Test_builtinDelete(t *testing.T) {
	var builtinDelete func(args ...nanojs.Object) (nanojs.Object, error)
	for _, f := range nanojs.GetAllBuiltinFunctions() {
		if f.Name == "delete" {
			builtinDelete = f.Value
			break
		}
	}
	if builtinDelete == nil {
		t.Fatal("builtin delete not found")
	}
	type args struct {
		args []nanojs.Object
	}
	tests := []struct {
		name      string
		args      args
		want      nanojs.Object
		wantErr   bool
		wantedErr error
		target    interface{}
	}{
		{name: "invalid-arg", args: args{[]nanojs.Object{&nanojs.String{},
			&nanojs.String{}}}, wantErr: true,
			wantedErr: nanojs.ErrInvalidArgumentType{
				Name:     "first",
				Expected: "map",
				Found:    "string"},
		},
		{name: "no-args",
			wantErr: true, wantedErr: nanojs.ErrWrongNumArguments},
		{name: "empty-args", args: args{[]nanojs.Object{}}, wantErr: true,
			wantedErr: nanojs.ErrWrongNumArguments,
		},
		{name: "3-args", args: args{[]nanojs.Object{
			(*nanojs.Map)(nil), (*nanojs.String)(nil), (*nanojs.String)(nil)}},
			wantErr: true, wantedErr: nanojs.ErrWrongNumArguments,
		},
		{name: "nil-map-empty-key",
			args: args{[]nanojs.Object{&nanojs.Map{}, &nanojs.String{}}},
			want: nanojs.UndefinedValue,
		},
		{name: "nil-map-nonstr-key",
			args: args{[]nanojs.Object{
				&nanojs.Map{}, &nanojs.Int{}}}, wantErr: true,
			wantedErr: nanojs.ErrInvalidArgumentType{
				Name: "second", Expected: "string", Found: "int"},
		},
		{name: "nil-map-no-key",
			args: args{[]nanojs.Object{&nanojs.Map{}}}, wantErr: true,
			wantedErr: nanojs.ErrWrongNumArguments,
		},
		{name: "map-missing-key",
			args: args{
				[]nanojs.Object{
					&nanojs.Map{Value: map[string]nanojs.Object{
						"key": &nanojs.String{Value: "value"},
					}},
					&nanojs.String{Value: "key1"}}},
			want: nanojs.UndefinedValue,
			target: &nanojs.Map{
				Value: map[string]nanojs.Object{
					"key": &nanojs.String{
						Value: "value"}}},
		},
		{name: "map-emptied",
			args: args{
				[]nanojs.Object{
					&nanojs.Map{Value: map[string]nanojs.Object{
						"key": &nanojs.String{Value: "value"},
					}},
					&nanojs.String{Value: "key"}}},
			want:   nanojs.UndefinedValue,
			target: &nanojs.Map{Value: map[string]nanojs.Object{}},
		},
		{name: "map-multi-keys",
			args: args{
				[]nanojs.Object{
					&nanojs.Map{Value: map[string]nanojs.Object{
						"key1": &nanojs.String{Value: "value1"},
						"key2": &nanojs.Int{Value: 10},
					}},
					&nanojs.String{Value: "key1"}}},
			want: nanojs.UndefinedValue,
			target: &nanojs.Map{Value: map[string]nanojs.Object{
				"key2": &nanojs.Int{Value: 10}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinDelete(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinDelete() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if tt.wantErr && !errors.Is(err, tt.wantedErr) {
				if err.Error() != tt.wantedErr.Error() {
					t.Errorf("builtinDelete() error = %v, wantedErr %v",
						err, tt.wantedErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("builtinDelete() = %v, want %v", got, tt.want)
				return
			}
			if !tt.wantErr && tt.target != nil {
				switch v := tt.args.args[0].(type) {
				case *nanojs.Map, *nanojs.Array:
					if !reflect.DeepEqual(tt.target, tt.args.args[0]) {
						t.Errorf("builtinDelete() objects are not equal "+
							"got: %+v, want: %+v", tt.args.args[0], tt.target)
					}
				default:
					t.Errorf("builtinDelete() unsuporrted arg[0] type %s",
						v.TypeName())
					return
				}
			}
		})
	}
}

func Test_builtinSplice(t *testing.T) {
	var builtinSplice func(args ...nanojs.Object) (nanojs.Object, error)
	for _, f := range nanojs.GetAllBuiltinFunctions() {
		if f.Name == "splice" {
			builtinSplice = f.Value
			break
		}
	}
	if builtinSplice == nil {
		t.Fatal("builtin splice not found")
	}
	tests := []struct {
		name      string
		args      []nanojs.Object
		deleted   nanojs.Object
		Array     *nanojs.Array
		wantErr   bool
		wantedErr error
	}{
		{name: "no args", args: []nanojs.Object{}, wantErr: true,
			wantedErr: nanojs.ErrWrongNumArguments,
		},
		{name: "invalid args", args: []nanojs.Object{&nanojs.Map{}},
			wantErr: true,
			wantedErr: nanojs.ErrInvalidArgumentType{
				Name: "first", Expected: "array", Found: "map"},
		},
		{name: "invalid args",
			args:    []nanojs.Object{&nanojs.Array{}, &nanojs.String{}},
			wantErr: true,
			wantedErr: nanojs.ErrInvalidArgumentType{
				Name: "second", Expected: "int", Found: "string"},
		},
		{name: "negative index",
			args:      []nanojs.Object{&nanojs.Array{}, &nanojs.Int{Value: -1}},
			wantErr:   true,
			wantedErr: nanojs.ErrIndexOutOfBounds},
		{name: "non int count",
			args: []nanojs.Object{
				&nanojs.Array{}, &nanojs.Int{Value: 0},
				&nanojs.String{Value: ""}},
			wantErr: true,
			wantedErr: nanojs.ErrInvalidArgumentType{
				Name: "third", Expected: "int", Found: "string"},
		},
		{name: "negative count",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: -1}},
			wantErr:   true,
			wantedErr: nanojs.ErrIndexOutOfBounds,
		},
		{name: "insert with zero count",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: 0},
				&nanojs.String{Value: "b"}},
			deleted: &nanojs.Array{Value: []nanojs.Object{}},
			Array: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.String{Value: "b"},
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2}}},
		},
		{name: "insert",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 0},
				&nanojs.String{Value: "c"},
				&nanojs.String{Value: "d"}},
			deleted: &nanojs.Array{Value: []nanojs.Object{}},
			Array: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 0},
				&nanojs.String{Value: "c"},
				&nanojs.String{Value: "d"},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2}}},
		},
		{name: "insert with zero count",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 0},
				&nanojs.String{Value: "c"},
				&nanojs.String{Value: "d"}},
			deleted: &nanojs.Array{Value: []nanojs.Object{}},
			Array: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 0},
				&nanojs.String{Value: "c"},
				&nanojs.String{Value: "d"},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2}}},
		},
		{name: "insert with delete",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 1},
				&nanojs.String{Value: "c"},
				&nanojs.String{Value: "d"}},
			deleted: &nanojs.Array{
				Value: []nanojs.Object{&nanojs.Int{Value: 1}}},
			Array: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 0},
				&nanojs.String{Value: "c"},
				&nanojs.String{Value: "d"},
				&nanojs.Int{Value: 2}}},
		},
		{name: "insert with delete multi",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2},
				&nanojs.String{Value: "c"},
				&nanojs.String{Value: "d"}},
			deleted: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2}}},
			Array: &nanojs.Array{
				Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.String{Value: "c"},
					&nanojs.String{Value: "d"}}},
		},
		{name: "delete all with positive count",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: 3}},
			deleted: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2}}},
			Array: &nanojs.Array{Value: []nanojs.Object{}},
		},
		{name: "delete all with big count",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: 5}},
			deleted: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2}}},
			Array: &nanojs.Array{Value: []nanojs.Object{}},
		},
		{name: "nothing2",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}}},
			Array: &nanojs.Array{Value: []nanojs.Object{}},
			deleted: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 0},
				&nanojs.Int{Value: 1},
				&nanojs.Int{Value: 2}}},
		},
		{name: "pop without count",
			args: []nanojs.Object{
				&nanojs.Array{Value: []nanojs.Object{
					&nanojs.Int{Value: 0},
					&nanojs.Int{Value: 1},
					&nanojs.Int{Value: 2}}},
				&nanojs.Int{Value: 2}},
			deleted: &nanojs.Array{Value: []nanojs.Object{&nanojs.Int{Value: 2}}},
			Array: &nanojs.Array{Value: []nanojs.Object{
				&nanojs.Int{Value: 0}, &nanojs.Int{Value: 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := builtinSplice(tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("builtinSplice() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.deleted) {
				t.Errorf("builtinSplice() = %v, want %v", got, tt.deleted)
			}
			if tt.wantErr && tt.wantedErr.Error() != err.Error() {
				t.Errorf("builtinSplice() error = %v, wantedErr %v",
					err, tt.wantedErr)
			}
			if tt.Array != nil && !reflect.DeepEqual(tt.Array, tt.args[0]) {
				t.Errorf("builtinSplice() arrays are not equal expected"+
					" %s, got %s", tt.Array, tt.args[0].(*nanojs.Array))
			}
		})
	}
}
