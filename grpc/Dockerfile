# Use the official Ubuntu base image
FROM ubuntu:24.10

WORKDIR app

# Update package list and install dependencies
RUN apt-get update && apt-get install -y \
    wget \
    build-essential \
    git

RUN wget https://go.dev/dl/go1.23.2.linux-arm64.tar.gz

RUN tar -C /usr/local -xzf go1.23.2.linux-arm64.tar.gz

RUN apt install -y protobuf-compiler

# Set the PATH environment variable
ENV PATH="/usr/local/go/bin:${PATH}"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="$PATH:$(go env GOPATH)/bin"

# Verify Go installation
RUN go version

