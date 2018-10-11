Install:
`yarn` or `npm i`

Execute tests:
`GOOS=js GOARCH=wasm go test -exec="node -r ./dom $(go env GOROOT)/misc/wasm/wasm_exec" -v`
