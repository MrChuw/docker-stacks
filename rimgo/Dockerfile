FROM --platform=$BUILDPLATFORM golang:alpine AS build

ARG TARGETARCH
WORKDIR /src

RUN apk --no-cache add ca-certificates git nodejs pnpm patch

COPY rimgo-src/ .

COPY custom_files/csp.go ./utils/csp.go

COPY patches/custom.patch /tmp/custom.patch
RUN patch -p1 < /tmp/custom.patch

RUN pnpm install && pnpm run build
RUN go mod download
RUN GOOS=linux GOARCH=$TARGETARCH CGO_ENABLED=0 go build -o rimgo-bin


FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/rimgo-bin ./rimgo

EXPOSE 3000

CMD ["./rimgo"]


