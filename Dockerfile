FROM golang:1.14-alpine as builder
ARG VERSION
COPY . /go/src
WORKDIR /go/src
RUN go build -mod vendor -o /go/bin/my-plugin -ldflags "-s -w -X main.version=${VERSION}"

FROM alpine:latest as run
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /bin
COPY --from=builder /go/bin/my-plugin /bin/my-plugin
ENTRYPOINT ["/bin/my-plugin"]
