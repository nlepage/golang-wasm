FROM nlepage/golang_wasm AS builder

COPY main.go src/hello/
RUN go install hello && mv bin/js_wasm/hello test.wasm

FROM nginx:1-alpine

COPY mime.types /etc/nginx/
COPY --from=builder /go/test.wasm /usr/local/go/misc/wasm/wasm_exec.* /usr/share/nginx/html/
