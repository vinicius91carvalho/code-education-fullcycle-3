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

# Set the PATH environment variable
ENV PATH="/usr/local/go/bin:${PATH}"

# Verify Go installation
RUN go version

