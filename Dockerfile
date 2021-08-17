FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o bookkeeper_mate .


FROM ttbb/bookkeeper:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/bookkeeper/mate

COPY --from=build /opt/sh/compile/pkg/bookkeeper_mate /opt/sh/bookkeeper/mate/bookkeeper_mate

COPY config/bk_server_original.conf /opt/sh/bookkeeper/conf/bk_server_original.conf

WORKDIR /opt/sh/bookkeeper

CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/bookkeeper/mate/scripts/start.sh"]