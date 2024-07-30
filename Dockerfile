FROM --platform=$BUILDPLATFORM golang:1.22.3-bookworm as vscode

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /service
COPY . /service

RUN  ln -sf /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

FROM --platform=$BUILDPLATFORM golang:1.22.3-bookworm as build

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /service
COPY . /service

RUN CGO_ENABLED=1 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app cmd/main.go

FROM --platform=$BUILDPLATFORM gcr.io/distroless/base-debian12:latest
COPY --chown=${USERNAME}:${GROUPNAME} --from=build /service/app /main

USER ${USERNAME}
ENTRYPOINT [ "/main" ]
