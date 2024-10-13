FROM golang:1.23.1-alpine AS builder

WORKDIR /usr/local/src

RUN apk --update upgrade && \
    apk add sqlite && \
    apk add --no-cache gcc musl-dev \
    rm -rf /var/cache/apk/*
# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 go build -o ./bin/app cmd/upstore/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /

CMD ["/app"]