FROM golang:1.22 as build
WORKDIR /storj-k6
ENV CGO_ENABLED=true
RUN --mount=type=cache,target=/root/.cache/go-build,id=gobuild \
    --mount=type=cache,target=/go/pkg/mod,id=gopkg \
    go install go.k6.io/xk6/cmd/xk6@latest
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build,id=gobuild \
    --mount=type=cache,target=/go/pkg/mod,id=gopkg \
    xk6 build --with github.com/elek/storj-k6=.

FROM alpine as image
RUN apk add --update micro bash
WORKDIR /root
COPY --from=build /storj-k6 /usr/local/bin
COPY examples/*.js .
COPY once.sh .
