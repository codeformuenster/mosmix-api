# stolen and remixed from https://github.com/simonw/timezones-api
FROM python:3.6-alpine3.7

ENV SQLITE_EXTENSIONS /usr/lib/mod_spatialite.so

RUN apk --repository http://nl.alpinelinux.org/alpine/edge/testing --no-cache add libspatialite-dev \
  && apk --no-cache add --virtual .builddeps build-base python3-dev \
  && pip install https://github.com/simonw/datasette/archive/446d47fdb005b3776bc06ad8d1f44b01fc2e938b.zip \
  && apk del .builddeps

COPY metadata.json /metadata.json

EXPOSE 8001

CMD ["datasette", "serve", "/mosmix/mosmix.spatialite", "--host", "0.0.0.0", "--port", "8001", "-m", "/metadata.json"]
