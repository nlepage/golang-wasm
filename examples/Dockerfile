FROM nlepage/golang_wasm AS builder

COPY hello/ src/hello/
COPY js-call/ src/js-call/
COPY go-call/ src/go-call/
RUN go install hello \
  && mv bin/js_wasm/hello hello.wasm \
  && go install js-call \
  && mv bin/js_wasm/js-call js-call.wasm \
  && go install go-call \
  && mv bin/js_wasm/go-call go-call.wasm

FROM nginx:1-alpine

COPY mime.types /etc/nginx/
COPY --from=builder /usr/local/go/misc/wasm/wasm_exec.js /go/*.wasm /usr/share/nginx/html/
COPY *.html /usr/share/nginx/html/
