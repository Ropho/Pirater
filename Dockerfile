# FROM golang:1.19-alpine

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# COPY ./bin/prog ./

# RUN apk add make
# RUN make

# CMD ["make", "run"]

FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN apk add make
RUN make

CMD ["make", "run"]