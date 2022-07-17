FROM node:14.19.3-alpine3.16 AS web-builder

WORKDIR /workspace/web

COPY web .

ENV NODE_OPTIONS --max_old_space_size=10240

RUN yarn install
RUN NEXT_TELEMETRY_DISABLED=1 yarn run export


FROM golang:1.18.3-alpine3.16 AS binary-builder

WORKDIR /workspace

ENV GOOS linux
ENV GOARCH amd64
ENV GO111MODULE on

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN env

COPY --from=web-builder "/workspace/web/dist" /workspace/cmd/borzoi/web

RUN apk --no-cache add ca-certificates gcc musl-dev

RUN go build -o borzoi cmd/borzoi/main.go


FROM alpine:3.13

WORKDIR /

RUN apk --no-cache add ca-certificates curl iproute2 tini

COPY --from=binary-builder "/workspace/borzoi" /
EXPOSE 8080

ENTRYPOINT ["/sbin/tini", "--", "/borzoi", "server"]

HEALTHCHECK --interval=5m --timeout=3s \
  CMD curl -f http://localhost:8080/ || exit 1