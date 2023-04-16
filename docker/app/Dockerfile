FROM alpine

WORKDIR /app

COPY ./bin/prog .
COPY ./config ./config
COPY ./docs ./docs

RUN apk add gcompat
RUN apk add tzdata

ENTRYPOINT ["./prog"]