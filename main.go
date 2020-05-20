package main

import (
	"fmt"
	"syscall/js"
)

var c chan bool

func init() {
	c = make(chan bool)
}

func add(this js.Value, i []js.Value) interface{} {
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	println(js.ValueOf(i[0].Int() + i[1].Int()).String())
	return js.ValueOf(i[0].Int() - i[1].Int())
}

func subtract(this js.Value, i []js.Value) interface{} {

	js.Global().Set("output", js.ValueOf(i[0].Int()-i[1].Int()))
	println(js.ValueOf("subtract!!").String())
	println(js.ValueOf(i[0].Int() - i[1].Int()).String())
	return js.ValueOf(i[0].Int() - i[1].Int())
}

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", message)
	document.Get("body").Call("appendChild", p)
	// c <- true
	return 1
}

func registerCallbacks() {

	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("printMessage", js.FuncOf(printMessage))
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	console_log := js.Global().Get("console").Get("log")
	console_log.Invoke("Hello wasm! invoke by go")

	js.Global().Call("eval", `
        console.log("hello, wasm! in console");
	`)

	registerCallbacks()

	js.Global().Set("println",
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			println("hello callback")
			println(args[0].String()) // Debug 语句 可以在浏览器调用的时候看到
			return nil
		}),
	)

	// printlnFn := js.Global().Get("println")
	// printlnFn.Invoke()
	// js.Global().Call("println", js.ValueOf("args!!!!"))

	select {}
	println("We are out of here")

}
