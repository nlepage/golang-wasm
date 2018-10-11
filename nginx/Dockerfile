FROM golang:1.11 AS builder

FROM nginx:1-alpine

COPY --from=builder /usr/local/go/misc/wasm/wasm_exec.* /usr/share/nginx/html/
COPY mime.types /etc/nginx/
