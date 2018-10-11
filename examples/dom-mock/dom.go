// +build js,wasm

package dom

import (
	"syscall/js"
)

func GetWindowLocation() string {
	return js.Global().Get("window").Get("location").String()
}

func GetDocumentTitle() string {
	return js.Global().Get("document").Get("title").String()
}
