package main

import (
	"fmt"

	"github.com/nlepage/golang-wasm/js/bind"
)

func main() {
	scope := &struct {
		Alert           func(string)  `js:"alert()"`
		Name            func() string `js:"helloName"`
		SetMessage      func(string)  `js:"helloMessage"`
		SetMessageTimes func(int)     `js:"helloMessageTimes"`
		LogMessage      func()        `js:"logHelloMessage()"`
	}{}
	bind.BindGlobals(scope)
	scope.Alert(fmt.Sprintf("Hello %s!", scope.Name()))
	scope.SetMessage("Hello from GOGOGOGO!!!")
	scope.SetMessageTimes(6)
	scope.LogMessage()
}
