FROM golang:1.23 AS builder
WORKDIR /usr/src/app
COPY hello.go .
RUN go build -v -o /usr/local/bin/app hello.go

FROM alpine:3.20
COPY --from=builder /usr/local/bin/app /usr/local/bin/app
CMD ["./usr/local/bin/app"]