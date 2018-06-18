package main

import (
	"fmt"
	"syscall/js"
)

var done = make(chan struct{})

func main() {
	callback := js.NewCallback(printMessage)
	defer callback.Close() // This is a good practice
	setPrintMessage := js.Global.Get("setPrintMessage")
	setPrintMessage.Invoke(callback)
	<-done
}

func printMessage(args []js.Value) {
	message := args[0].String()
	fmt.Println(message)
	done <- struct{}{} // Notify printMessage has been called
}
