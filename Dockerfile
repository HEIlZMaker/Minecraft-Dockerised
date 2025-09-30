FROM golang:alpine3.22 AS builder 
WORKDIR /build
COPY ./ ./
RUN go build main.go

FROM gcr.io/distroless/java21-debian12
WORKDIR /bin
COPY --from=builder /build/main ./
WORKDIR /mc
VOLUME [ "/mc" ]
EXPOSE 25565/tcp
ENTRYPOINT [ "main" ]