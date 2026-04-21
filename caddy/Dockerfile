ARG CADDY_VERSION=2.11.2
FROM caddy:${CADDY_VERSION}-builder AS builder

RUN xcaddy build \
    --with github.com/lucaslorentz/caddy-docker-proxy/v2 \
    --with github.com/mholt/caddy-l4/layer4 \
    --with github.com/caddy-dns/cloudflare \
    --with github.com/WeidiDeng/caddy-cloudflare-ip \
    --with github.com/fvbommel/caddy-combine-ip-ranges

FROM caddy:${CADDY_VERSION}-alpine

RUN apk add --no-cache nss-tools ca-certificates

COPY --from=builder /usr/bin/caddy /usr/bin/caddy

CMD ["caddy", "docker-proxy"]
#CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--watch"]