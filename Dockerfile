FROM golang:1.10-alpine as builder

ENV CADDYPATH=$GOPATH/src/github.com/mholt/caddy

RUN apk add --no-cache git gcc musl-dev

COPY caddy-mosmix-prest/mosmixapi.go ${GOPATH}/src/github.com/codeformuenster/mosmix-api/caddy-mosmix-prest/mosmixapi.go

# fetch deps
RUN go get ./...
RUN go get github.com/caddyserver/builds github.com/captncraig/cors/caddy

# Insert moxmix-prest plugin
RUN sed -i '/\(imported\)/a_ "github.com/codeformuenster/mosmix-api/caddy-mosmix-prest"\n_ "github.com/captncraig/cors/caddy"' "${CADDYPATH}/caddy/caddymain/run.go"
RUN sed -i '/"browse"/i"mosmixapi",' "${CADDYPATH}/caddyhttp/httpserver/plugin.go"

# Build the binary
RUN cd "${CADDYPATH}/caddy" \
    && GOOS=linux GOARCH=amd64 go run build.go -goos=$GOOS -goarch=$GOARCH \
    && mkdir -p /install \
    && mv caddy /install
FROM swaggerapi/swagger-ui:3.17.6 as swaggerui

WORKDIR /usr/share/nginx/html

COPY swagger/* ./

RUN (head -n 38 index.html; cat swagger.js; tail -n 3 index.html) > index.html.tmp && \
  mv index.html.tmp index.html

FROM alpine:3.8

RUN apk add --no-cache openssh-client ca-certificates

COPY --from=builder /install/caddy /usr/bin/caddy

COPY Caddyfile /etc/Caddyfile
COPY queries /srv/queries
COPY --from=swaggerui /usr/share/nginx/html /srv/html

WORKDIR /srv

ENTRYPOINT ["/usr/bin/caddy"]
CMD ["--conf=/etc/Caddyfile", "--log=stdout", "--agree=true", "--root=/srv"]
