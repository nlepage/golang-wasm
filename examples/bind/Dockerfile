FROM nlepage/golang_wasm AS builder

COPY --from=nlepage/golang_wasm:js-bind /go/src /go/src
COPY ./ src/bind/
RUN go build -o test.wasm bind

FROM nlepage/golang_wasm:nginx

COPY wasm_exec.html /usr/share/nginx/html/
COPY --from=builder /go/test.wasm /usr/share/nginx/html/
