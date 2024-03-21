FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder
LABEL maintainer="nekohasekai <contact-git@sekai.icu>"
COPY . /go/src/github.com/ebrahimtahernejad/sing-box-warp
WORKDIR /go/src/github.com/ebrahimtahernejad/sing-box-warp
ARG TARGETOS TARGETARCH
ARG GOPROXY=""
ENV GOPROXY ${GOPROXY}
ENV CGO_ENABLED=0
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH
RUN set -ex \
    && apk add git build-base \
    && export COMMIT=$(git rev-parse --short HEAD) \
    && export VERSION=$(go run ./cmd/internal/read_tag) \
    && go build -v -trimpath -tags \
        "with_gvisor,with_quic,with_dhcp,with_wireguard,with_ech,with_utls,with_reality_server,with_acme,with_clash_api" \
        -o /go/bin/sing-box \
        -ldflags "-X \"github.com/ebrahimtahernejad/sing-box-warp/constant.Version=$VERSION\" -s -w -buildid=" \
        ./cmd/sing-box
FROM --platform=$TARGETPLATFORM alpine AS dist
LABEL maintainer="nekohasekai <contact-git@sekai.icu>"
RUN set -ex \
    && apk upgrade \
    && apk add bash tzdata ca-certificates \
    && rm -rf /var/cache/apk/*
COPY --from=builder /go/bin/sing-box /usr/local/bin/sing-box
ENTRYPOINT ["sing-box"]
