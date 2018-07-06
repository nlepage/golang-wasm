FROM nlepage/golang_wasm:hello AS hello
FROM nlepage/golang_wasm:js-call AS js-call
FROM nlepage/golang_wasm:go-call AS go-call
FROM nlepage/golang_wasm:long-running AS long-running
FROM nlepage/golang_wasm:bind-counter AS bind-counter

FROM nlepage/golang_wasm:nginx

COPY --from=hello /usr/share/nginx/html/ /usr/share/nginx/html/hello/
COPY --from=js-call /usr/share/nginx/html/ /usr/share/nginx/html/js-call/
COPY --from=go-call /usr/share/nginx/html/ /usr/share/nginx/html/go-call/
COPY --from=long-running /usr/share/nginx/html/ /usr/share/nginx/html/long-running/
COPY --from=bind-counter /usr/share/nginx/html/ /usr/share/nginx/html/bind-counter/
COPY index.html /usr/share/nginx/html/
