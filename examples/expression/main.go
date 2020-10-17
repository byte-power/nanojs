package main

import (
	"fmt"
	"sync"

	"github.com/zeaphoo/nanojs/v2"
)

const expr1 = `
player.item.car += 2
`
const expr2 = `
player.item.car -= 1
`
const exprCondition = `
// comment line
/*
comment block
// line in block
*/
console.log(1 >= 0, true && true, true || false, !false)
[
	"logic",
	!(1>1)&&true||false,
	true&&-1!=1,
	!"" == ![],
	!1,

	"int vs float",
	["0.1+0.2=", 0.1+0.2],
	["4*5/3+1=", 4*5/3+1],
	["4*5/3.0+1=", 4*5/3.0+1],
	((1+2)*3)/0.3333,

	"bitwise",
	17^3,
	int(((1+2)*3)/0.5)<<1
]
`
const exprFinish = `
player.item.car>0
`
const (
	globalConsole = "console"
	globalPlayer  = "player"
)

var (
	globalConsoleFuncs = map[string]nanojs.Object{
		"log": &nanojs.UserFunction{Value: func(args ...nanojs.Object) (ret nanojs.Object, err error) {
			fmt.Println(args)
			return
		}},
	}
)

func main() {
	moduleMap := nanojs.NewModuleMap()
	player := newPlayerModule("p0")

	scripts := map[string]*nanojs.Expression{
		"_1": nanojs.NewExpression([]byte(expr1)),
		"_2": nanojs.NewExpression([]byte(expr2)),

		"condition": nanojs.NewExpression([]byte(exprCondition)),
		"finish":    nanojs.NewExpression([]byte(exprFinish)),
	}

	compileds := map[string]*nanojs.CompiledExpression{}
	for _, script := range scripts {
		script.SetImports(moduleMap)
		script.Add(globalConsole, &nanojs.Map{Value: globalConsoleFuncs})
		script.Add(globalPlayer, player)
	}
	for i, script := range scripts {
		fmt.Println("compiling no.", i)
		compiled, err := script.Compile()
		if err != nil {
			panic(err)
		}
		compileds[i] = compiled
	}

	player = newPlayerModule("p1")
	for i, compiled := range compileds {
		if i[0] != '_' {
			continue
		}
		fmt.Println("\nrunning no.", i)
		cloned := compiled.Clone()
		cloned.Set(globalPlayer, player)
		_, err := cloned.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
	runCompiled := func(player *PlayerModule, name string) {
		fmt.Println("\nrunning", name)
		cloned := compileds[name]
		cloned.Set(globalPlayer, player)
		ret, err := cloned.Run()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ret)
	}
	runCompiled(player, "condition")
	runCompiled(player, "finish")
}

type Item struct {
	count int64
}

type PlayerModule struct {
	nanojs.ObjectImpl

	playerID   string
	attributes map[string]interface{}
	items      map[string]Item

	lock sync.Mutex
}

func newPlayerModule(id string) *PlayerModule {
	mod := &PlayerModule{
		playerID: id,
		items:    make(map[string]Item),
	}
	return mod
}

func (player *PlayerModule) Typename() string {
	return "player"
}

func (player *PlayerModule) String() string {
	return player.playerID
}

func (player *PlayerModule) IndexGet(arg nanojs.Object) (ret nanojs.Object, err error) {
	key, _ := arg.(*nanojs.String)
	switch key.Value {
	case "item":
		ret = &ItemModule{player: player}
	case "verbose":
		ret = &nanojs.UserFunction{Value: func(args ...nanojs.Object) (ret nanojs.Object, err error) {
			fmt.Println(player.items)
			return
		}}
	}
	fmt.Println("PlayerModule.IndexGet:", key, ret)
	return
}

func (player *PlayerModule) changeItemCount(name string, count int64) {
	player.lock.Lock()
	defer player.lock.Unlock()
	item := player.items[name]
	item.count += count
	player.items[name] = item
}

func (player *PlayerModule) setItemCount(name string, count int64) {
	player.lock.Lock()
	defer player.lock.Unlock()
	item := player.items[name]
	item.count = count
	player.items[name] = item
}

type ItemModule struct {
	nanojs.ObjectImpl

	player   *PlayerModule
	itemName string
}

func (it *ItemModule) Typename() string {
	return "item"
}

func (it *ItemModule) String() string {
	return it.itemName
}

func (it *ItemModule) Copy() nanojs.Object {
	return &ItemModule{
		itemName: it.itemName,
		player:   it.player,
	}
}

func (it *ItemModule) IndexGet(arg nanojs.Object) (ret nanojs.Object, err error) {
	key, _ := nanojs.ToString(arg)
	if key != "" {
		ret = &nanojs.Int{Value: it.player.items[key].count}
	}
	fmt.Println("ItemModule.IndexGet:", key, ret)
	return
}

// IndexSet sets the value for the given key.
func (it *ItemModule) IndexSet(index, value nanojs.Object) (err error) {
	key, ok := nanojs.ToString(index)
	if !ok {
		err = nanojs.ErrInvalidIndexType
		return
	}
	num, ok := nanojs.ToInt64(value)
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Expected: "number",
			Found:    value.TypeName(),
		}
		return
	}
	it.player.setItemCount(key, num)
	fmt.Println("ItemModule.IndexSet:", key, num)
	return nil
}

func (it *ItemModule) add(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) == 0 {
		return
	}
	arg, _ := args[0].(*nanojs.Int)
	if arg == nil || arg.Value == 0 {
		return
	}
	it.player.changeItemCount(it.itemName, arg.Value)
	return
}

func (it *ItemModule) sub(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) == 0 {
		return
	}
	arg, _ := args[0].(*nanojs.Int)
	if arg == nil || arg.Value == 0 {
		return
	}
	it.player.changeItemCount(it.itemName, -arg.Value)
	return
}
