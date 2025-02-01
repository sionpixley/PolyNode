FROM ubuntu:latest AS base

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y ca-certificates xz-utils

FROM golang:1.23.5-alpine AS build

WORKDIR /PolyNode

COPY . .

RUN go build -ldflags="-s -w" -tags=prod -o polyn ./cmd/polyn
RUN cd install && go build -ldflags="-s -w" -o ../setup ./cmd/setup && cd ..
RUN cd uninstall && go build -ldflags="-s -w" -o ../uninstall-linux && cd ..

FROM base

WORKDIR /PolyNode

COPY --from=build /PolyNode/setup .
RUN mkdir PolyNode
COPY --from=build /PolyNode/polyn ./PolyNode/polyn
RUN cd PolyNode && mkdir uninstall && cd ..
COPY --from=build /PolyNode/uninstall-linux ./PolyNode/uninstall/uninstall-linux
RUN cd ./PolyNode/uninstall && mv uninstall-linux uninstall && cd ..

ENV SHELL=/bin/bash
ENV PATH=$PATH:/root/.PolyNode:/root/.PolyNode/nodejs/bin
RUN ./setup

CMD [ "sleep", "infinity" ]
