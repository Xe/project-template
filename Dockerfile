ARG GO_VERSION=1.24
ARG ALPINE_VERSION=edge
FROM --platform=${BUILDPLATFORM} golang:${GO_VERSION}-alpine AS build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN --mount=type=cache,target=/root/.cache GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -gcflags "all=-N -l" -o /app/bin/web ./cmd/web

FROM alpine:${ALPINE_VERSION} AS run
WORKDIR /app

ARG COMPONENT

RUN apk -U add ca-certificates mailcap

COPY --from=build /app/bin/web /app/bin/web

CMD ["/app/bin/web"]

LABEL org.opencontainers.image.source="https://git.xeserv.us/xe/project-template"