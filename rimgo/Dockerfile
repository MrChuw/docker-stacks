FROM --platform=$BUILDPLATFORM golang:alpine AS build

ARG TARGETARCH
WORKDIR /src

# Instala dependências de build + patch
RUN apk --no-cache add ca-certificates git nodejs pnpm patch

# 1. Copia o código do submódulo
COPY rimgo-src/ .

# 2. Injeta seus arquivos customizados (antes de aplicar o patch)
# Certifique-se que o caminho 'utils/' existe no rimgo
COPY custom_files/csp.go ./utils/csp.go

# 3. Aplica o seu patch
COPY patches/custom.patch /tmp/custom.patch
RUN patch -p1 < /tmp/custom.patch

# 4. Build padrão do rimgo
RUN pnpm install && pnpm run build
RUN go mod download
RUN GOOS=linux GOARCH=$TARGETARCH CGO_ENABLED=0 go build -o rimgo-bin

# --- Imagem Final ---
FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates

# Copia o binário e os assets necessários
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/rimgo-bin ./rimgo
# COPY --from=build /src/static ./static
# COPY --from=build /src/views ./views

EXPOSE 3000

CMD ["./rimgo"]


