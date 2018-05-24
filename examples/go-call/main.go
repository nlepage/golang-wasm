package main

import (
	"fmt"
	"syscall/js"
)

var (
	no = 1
)

func main() {
	setPrintln := js.Global.Get("setPrintMessage")
	printlnCallback := js.NewCallback(func(args []js.Value) {
		message := args[0].String()
		fmt.Printf("Message no %d printed by Go callback: %s\n", no, message)
		no++
	})
	setPrintln.Invoke(printlnCallback)
	select {} // Use select to allow for callback goroutines to live on
}
