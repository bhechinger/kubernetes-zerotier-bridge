FROM golang:1.15-alpine AS go-build
RUN mkdir /app
ADD manage/ /app/
WORKDIR /app

RUN CGO_ENABLED=0 go build -ldflags "-w"

FROM alpine
MAINTAINER Juan Manuel Vera - verajm@gmail.com

ENV ZEROTIER_VERSION=1.6.4

RUN set -eux; \
    apk add --no-cache \
      libgcc \
      libstdc++ \
    ; \
    apk add --no-cache --virtual build-dependencies \
      build-base \
      linux-headers \
    ; \
    apk add --update supervisor \
        bash \
        iptables\
        openrc \
        curl \
    ;\
    wget https://github.com/zerotier/ZeroTierOne/archive/$ZEROTIER_VERSION.zip -O /zerotier.zip; \
    unzip /zerotier.zip -d /; \
    cd /ZeroTierOne-$ZEROTIER_VERSION; \
    make; \
    DESTDIR=/tmp/build make install; \
    mv /tmp/build/usr/sbin/* /usr/sbin/; \
    mkdir /var/lib/zerotier-one; \
    apk del build-dependencies; \
    rm -rf /tmp/build; \
    rm -rf /ZeroTierOne-$ZEROTIER_VERSION; \
    rm -rf /zerotier.zip; \
    zerotier-one -v; \
    rc-update add iptables ;\
    echo "tun" >> /etc/modules

COPY files/supervisor-zerotier.conf /etc/supervisor/supervisord.conf
COPY files/entrypoint.sh /entrypoint.sh
COPY --from=go-build /app/manage /usr/sbin/manage

RUN chmod 755 /entrypoint.sh
RUN chmod 755 /usr/sbin/manage

VOLUME ["/var/lib/zerotier-one"]
ENTRYPOINT ["/entrypoint.sh"]