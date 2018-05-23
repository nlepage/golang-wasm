### Hello example

Run with:

```sh
docker container run -dP nlepage/golang_wasm:hello

# Find out which host port is used
docker container ps
```

Visit http://localhost:32XXX/wasm_exec.html, open browser console and click Run button
