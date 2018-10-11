FROM golang:1.11 AS builder

COPY ./ src/js-call/
RUN GOOS=js GOARCH=wasm go build -o test.wasm js-call

FROM nlepage/golang_wasm:nginx

COPY --from=builder /go/test.wasm /usr/share/nginx/html/
