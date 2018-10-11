// +build js,wasm

package dom

import (
	"testing"
)

func TestGetWindowLocation(t *testing.T) {
	const expected = "http://localhost:8080/wasm_exec.html"
	if actual := GetWindowLocation(); actual != expected {
		t.Errorf("%s != %s\n", actual, expected)
	}
}

func TestGetDocumentTitle(t *testing.T) {
	const expected = "Go wasm"
	if actual := GetDocumentTitle(); actual != expected {
		t.Errorf("%s != %s\n", actual, expected)
	}
}
