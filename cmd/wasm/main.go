package main

import (
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("sub", js.FuncOf(sub))
	js.Global().Set("mul", js.FuncOf(mul))
	js.Global().Set("div", js.FuncOf(div))
}

func add(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() + i[1].Int())
	return setResult(result)
}

func sub(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() - i[1].Int())
	return setResult(result)
}

func mul(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() * i[1].Int())
	return setResult(result)
}

func div(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() / i[1].Int())
	return setResult(result)
}

func setResult(val js.Value) error {
	js.Global().Get("document").Call("getElementById", "output").Set("value", val)
	return nil
}

func main() {
	c := make(chan struct{}, 0)
	println("WASM Go calculator initialized")
	registerCallbacks()
	<-c
}
