FROM golang:1.10-alpine as build-stage

RUN apk add --no-cache git

WORKDIR /go/src/github.com/prest/prest

RUN git clone https://github.com/prest/prest.git . && \
  git reset --hard eb190b4458e76b635b217d202696469b5c6e76b5

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags netgo -ldflags='-s -w -extldflags -static'
RUN go build -ldflags "-s -w"

FROM alpine:3.7

COPY --from=build-stage /go/src/github.com/prest/prest/prest /

COPY api/prest.toml /prest.toml
COPY api/queries /prest-queries

CMD ["/prest"]
