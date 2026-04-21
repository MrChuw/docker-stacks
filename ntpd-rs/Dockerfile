FROM alpine:edge

RUN apk add --no-cache \
    --repository=https://dl-cdn.alpinelinux.org/alpine/edge/testing \
    ntpd-rs

# Garante permissão total no /run para o socket
RUN mkdir -p /run && chmod 777 /run

# Links simbólicos
RUN DAEMON_PATH=$(find /usr -name "*ntp-daemon*") && ln -s $DAEMON_PATH /usr/local/bin/ntp-daemon && \
    METRICS_PATH=$(find /usr -name "*ntp-metrics-exporter*") && ln -s $METRICS_PATH /usr/local/bin/ntp-metrics-exporter && \
    CTL_PATH=$(find /usr -name "*ntp-ctl*") && ln -s $CTL_PATH /usr/local/bin/ntp-ctl

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]