FROM golang:1.10-alpine as build-stage

RUN apk add --no-cache git

WORKDIR /go/src/github.com/prest/prest

RUN git clone https://github.com/prest/prest.git . && \
  git reset --hard 3f7e7d98726602c62a27a422092ffe003400ff47

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags netgo -ldflags='-s -w -extldflags -static'
RUN go build -ldflags "-s -w"

FROM alpine:3.8

COPY --from=build-stage /go/src/github.com/prest/prest/prest /

COPY ./prest.toml /prest.toml
COPY ./queries /prest-queries

CMD ["/prest"]
