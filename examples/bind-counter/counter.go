package main

import (
	"fmt"

	"github.com/nlepage/golang-wasm/js/bind"
)

type Example struct {
	Alert      func(string) `js:"alert()"`
	Counter    func() int   `js:"counter"`
	SetCounter func(int)    `js:"counter"`
}

func main() {
	example := &Example{}

	if err := bind.BindGlobals(example); err != nil {
		panic(err)
	}

	example.SetCounter(example.Counter() + 1)
	example.Alert(fmt.Sprintf("Run no %d", example.Counter()))
}
