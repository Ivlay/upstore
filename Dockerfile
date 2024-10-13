# =============================================================================
#  Multi-stage Dockerfile Example
# =============================================================================
#  This is a simple Dockerfile that will build an image of scratch-base image.
#  Usage:
#    docker build -t simple:local . && docker run --rm simple:local
# =============================================================================

# -----------------------------------------------------------------------------
#  Build Stage
# -----------------------------------------------------------------------------
  FROM golang:alpine3.18 AS build

  # Important:
  #   Because this is a CGO enabled package, you are required to set it as 1.
  ENV CGO_ENABLED=1
  
  RUN apk add --no-cache \
      # Important: required for go-sqlite3
      gcc \
      # Required for Alpine
      musl-dev
  
  WORKDIR /workspace
  
  COPY . /workspace/
  
  RUN \
      cd cmd/upstore && \
      go mod init github.com/mattn/sample && \
      go mod edit -replace=github.com/mattn/go-sqlite3=../.. && \
      go mod tidy && \
      go install -ldflags='-s -w -extldflags "-static"' ./main.go
  
  RUN \
      # Smoke test
      set -o pipefail; \
      /go/bin/upstore | grep 99\ こんにちは世界099
  
  # -----------------------------------------------------------------------------
  #  Main Stage
  # -----------------------------------------------------------------------------
  FROM scratch
  
  COPY --from=build /go/bin/upstore /usr/local/bin/upstore
  
  ENTRYPOINT [ "/usr/local/bin/upstore" ]