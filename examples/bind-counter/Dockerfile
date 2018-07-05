FROM nlepage/golang_wasm AS builder

COPY ./ src/bind-counter/
RUN go get github.com/nlepage/golang-wasm/js/bind
RUN go build -o test.wasm bind-counter

FROM nlepage/golang_wasm:nginx

COPY wasm_exec.html /usr/share/nginx/html/
COPY --from=builder /go/test.wasm /usr/share/nginx/html/
