FROM nlepage/golang_wasm:hello AS hello
FROM nlepage/golang_wasm:js-call AS js-call
FROM nlepage/golang_wasm:go-call AS go-call

FROM nlepage/golang_wasm:nginx

COPY --from=hello /usr/share/nginx/html/ /usr/share/nginx/html/hello/
COPY --from=js-call /usr/share/nginx/html/ /usr/share/nginx/html/js-call/
COPY --from=go-call /usr/share/nginx/html/ /usr/share/nginx/html/go-call/
COPY index.html /usr/share/nginx/html/
