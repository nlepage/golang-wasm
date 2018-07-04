package main

import (
	"fmt"

	"github.com/nlepage/golang-wasm/js/bind"
)

type example struct {
	I func() int
	F func() float64
	B func() bool
	S func() string
}

func main() {
	scope := &struct {
		Alert           func(string)   `js:"alert()"`
		Name            func() string  `js:"helloName"`
		SetMessage      func(string)   `js:"helloMessage"`
		SetMessageTimes func(int)      `js:"helloMessageTimes"`
		LogMessage      func()         `js:"logHelloMessage()"`
		GetName         func() string  `js:"getName()"`
		Obj             func() example `js:"obj"`
	}{}
	bind.BindGlobals(scope)
	scope.Alert(fmt.Sprintf("Hello %s!", scope.Name()))
	scope.SetMessage("Hello from GOGOGOGO!!!")
	scope.SetMessageTimes(6)
	scope.LogMessage()
	fmt.Println(scope.GetName())
	obj := scope.Obj()
	fmt.Println(obj.F(), obj.I(), obj.B(), obj.S())
}
