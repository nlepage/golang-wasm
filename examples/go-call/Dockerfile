FROM nlepage/golang_wasm AS builder

COPY ./ src/go-call/
RUN go build -o test.wasm go-call

FROM nlepage/golang_wasm:nginx

COPY wasm_exec.html /usr/share/nginx/html/
COPY --from=builder /go/test.wasm /usr/share/nginx/html/
