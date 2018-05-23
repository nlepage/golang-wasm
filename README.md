Compile:
```sh
GOOS=js GOARCH=wasm <PATH_TO_SDK>/bin/go build -o test.wasm
```

Serve:
```sh
docker container run -v $(pwd):/usr/share/nginx/html:ro -v $(pwd)/mime.types:/etc/nginx/mime.types:ro -d -P nginx:1-alpine
```
