package main

import (
	"syscall/js"
)

func main() {
	alert := js.Global.Get("alert")
	alert.Invoke("Hello Go wasm !")
}
