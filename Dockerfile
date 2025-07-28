FROM golang:alpine3.22 AS builder
WORKDIR /build
COPY /home/heilz/dev/docker/distroless-minecraft/godev/main.go ./
RUN go build envvars.go

FROM gcr.io/distroless/java21-debian12
WORKDIR /bin
COPY --from=builder /build/main ./
RUN /bin/main
WORKDIR /world
VOLUME [ "/world" ]