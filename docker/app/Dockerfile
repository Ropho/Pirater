FROM jrottenberg/ffmpeg:alpine

WORKDIR /app

COPY ./bin/prog .
COPY ./config ./config

RUN apk add gcompat
RUN apk add tzdata

ENTRYPOINT ["./prog"]