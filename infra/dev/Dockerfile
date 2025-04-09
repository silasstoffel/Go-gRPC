FROM golang:1.24 AS builder

ARG PROTOC_VERSION=21.12

# Install protoc
RUN apt-get update && apt-get install -y unzip curl && \
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local && \
    rm protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# https://grpc.io/docs/languages/go/quickstart/
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/ktr0731/evans@latest

RUN export PATH="$PATH:$(go env GOPATH)/bin"


WORKDIR /app

COPY . .


CMD ["tail", "-f", "/dev/null"]

#RUN go mod download
#RUN go build -o app .

#FROM debian:bookworm-slim

#ARG PROTOC_VERSION=21.12
#RUN apt-get update && apt-get install -y curl unzip && \
#    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
#    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local && \
#    rm protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
#    apt-get clean && rm -rf /var/lib/apt/lists/*

#WORKDIR /app

# Copia o binário da fase de build
#COPY --from=builder /app/app .

# Comando de execução
#CMD ["./app"]
