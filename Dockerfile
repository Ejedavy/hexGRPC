FROM golang:1.19.3 as builder
WORKDIR /go/src/hex-grpc
COPY . .
RUN go build -o bin/server cmd/main.go

FROM alpine:latest as production
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/hex-grpc/bin/server /bin/server
EXPOSE 9000
ENTRYPOINT [ "/bin/server" ]