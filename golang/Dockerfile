FROM golang:1.10 AS builder

WORKDIR /tmp
ENV USER=go
RUN git clone https://github.com/neelance/go && cd go && git checkout wasm-wip
RUN cd go/src && ./all.bash
RUN mv go/VERSION.cache go/VERSION && rm -rf go/.git* go/pkg/obj

FROM buildpack-deps:stretch-scm

RUN apt-get update && apt-get install -y --no-install-recommends \
    g++ \
    gcc \
    libc6-dev \
    make \
    pkg-config \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /tmp/go/ /usr/local/go/

ENV GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:$PATH \
    GOARCH=wasm \
    GOOS=js

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
