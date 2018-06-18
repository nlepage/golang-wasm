FROM nlepage/golang_wasm AS builder

COPY ./ src/hello/
WORKDIR /go/src/hello
RUN go build -o test.wasm helloworld.go

FROM nlepage/golang_wasm:nginx

COPY --from=builder /go/src/hello/test.wasm /usr/share/nginx/html/
