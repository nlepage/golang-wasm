FROM golang:1.11 AS builder

COPY ./ src/long-running/
RUN GOOS=js GOARCH=wasm go build -o test.wasm long-running

FROM nlepage/golang_wasm:nginx

COPY wasm_exec.html /usr/share/nginx/html/
COPY --from=builder /go/test.wasm /usr/share/nginx/html/
