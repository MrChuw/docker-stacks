# Builds Request Baskets service using multi-stage builds
# Version 1.3 - Multi-arch tailored

# Stage 1. Building service
FROM --platform=$BUILDPLATFORM golang:latest AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /go/src/rbaskets
COPY request-baskets/ .

RUN GIT_VERSION="$(git describe --dirty='*' || git symbolic-ref -q --short HEAD 2>/dev/null || echo "unknown")" \
  && GIT_COMMIT="$(git rev-parse HEAD 2>/dev/null || echo "unknown")" \
  && GIT_COMMIT_SHORT="$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")" \
  && set -x \
  && CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -installsuffix cgo \
    -ldflags="-w -s -X main.GitVersion=${GIT_VERSION} -X main.GitCommit=${GIT_COMMIT} -X main.GitCommitShort=${GIT_COMMIT_SHORT}" \
    -o /go/bin/rbaskets

# Stage 2. Packaging into alpine
FROM alpine:latest
LABEL maintainer="vladimir_l@gmx.net"
RUN apk --no-cache add ca-certificates
VOLUME /var/lib/rbaskets

COPY request-baskets/docker/entrypoint.sh /bin/entrypoint
COPY --from=builder /go/bin/rbaskets /bin/rbaskets

EXPOSE 55555
CMD ["/bin/entrypoint"]
