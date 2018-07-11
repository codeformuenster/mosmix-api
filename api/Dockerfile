FROM golang:1.10-alpine as build-stage

WORKDIR /go/src/github.com/prest/prest

COPY . .

RUN go build -ldflags "-s -w -extldflags -static"

FROM alpine:3.8

COPY --from=build-stage /go/src/github.com/prest/prest/prest /

COPY ./prest.toml /prest.toml
COPY ./queries /prest-queries

CMD ["/prest"]
