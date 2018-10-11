FROM golang:1.11 AS builder

COPY ./ src/hello/
RUN GOOS=js GOARCH=wasm go build -o test.wasm hello

FROM nlepage/golang_wasm:nginx

COPY --from=builder /go/test.wasm /usr/share/nginx/html/
