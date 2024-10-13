FROM golang:1.23.1 AS builder

WORKDIR /app

RUN apt-get update && \
apt-get -y install gcc && \
rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

ENV CGO_ENABLED=1

RUN CGO_ENABLED=1 GOOS=linux go build -o /.bin/app

FROM scratch

WORKDIR /

COPY --from=builder /.bin/app /.bin/app

CMD ["/.bin/app"]