FROM --platform=$BUILDPLATFORM node:18-alpine AS frontend
WORKDIR /src/web
# Copia o frontend de dentro da subpasta rustlog-src
COPY rustlog-src/web/ .
RUN yarn install --ignore-optional
RUN yarn build

FROM --platform=$BUILDPLATFORM rust:1.95-bookworm AS chef
USER root
ENV CARGO_PROFILE_RELEASE_LTO=true
RUN cargo install cargo-chef
WORKDIR /app

FROM --platform=$BUILDPLATFORM chef AS planner
COPY . . 
RUN apt-get update && apt-get install -y patch
RUN patch -p1 -d rustlog-src < patches/first.patch
WORKDIR /app/rustlog-src
RUN cargo chef prepare --recipe-path ../recipe.json

FROM --platform=$BUILDPLATFORM chef AS builder

ARG TARGETPLATFORM
RUN case "${TARGETPLATFORM}" in \
      "linux/arm64") echo "aarch64-unknown-linux-gnu" > /target.txt && echo "-C linker=aarch64-linux-gnu-gcc" > /flags.txt ;; \
      "linux/amd64") echo "x86_64-unknown-linux-gnu" > /target.txt && echo "-C linker=x86_64-linux-gnu-gcc" > /flags.txt ;; \
      *) exit 1 ;; \
    esac

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update && \
    apt-get install -yq build-essential g++-aarch64-linux-gnu binutils-aarch64-linux-gnu patch

RUN rustup target add "$(cat /target.txt)"

COPY --from=planner /app/recipe.json recipe.json
RUN RUSTFLAGS="$(cat /flags.txt)" cargo chef cook --target "$(cat /target.txt)" --release --recipe-path recipe.json

COPY . .
RUN patch -p1 -d rustlog-src < patches/first.patch
COPY --from=frontend /src/web rustlog-src/web/

WORKDIR /app/rustlog-src
RUN RUSTFLAGS="$(cat /flags.txt)" cargo build --target "$(cat /target.txt)" --release
RUN TARGET_DIR=$(cat /target.txt) && \
    mv "target/${TARGET_DIR}/release/rustlog" /output

FROM debian:bookworm-slim AS runtime
RUN apt-get update && apt-get install -y gettext-base ca-certificates && rm -rf /var/lib/apt/lists/*
RUN useradd rustlog && mkdir /logs && chown rustlog: /logs

COPY --from=builder /output /usr/local/bin/rustlog
RUN chmod +x /usr/local/bin/rustlog

USER rustlog
ENV RUST_BACKTRACE=1
CMD ["/usr/local/bin/rustlog"]