FROM golang:1.17.0 AS builder

LABEL stage=server-intermediate

ARG GOPATH=/go/src
ENV GOPATH=${GOPATH}

WORKDIR /go/src/github.com/infobloxopen/protoc-gen-preprocess
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/usr/bin/protoc-gen-preprocess main.go

RUN ls -lR /go/bin

FROM infoblox/atlas-gentool:latest AS runner

COPY --from=builder /out/usr/bin/protoc-gen-preprocess /usr/bin/protoc-gen-preprocess
COPY --from=builder /go/bin/* /usr/bin/
COPY --from=builder /go/src/github.com/infobloxopen/protoc-gen-preprocess/options/*.proto /go/src/github.com/infobloxopen/protoc-gen-preprocess/options
# COPY --from=builder /go/src/github.com/infobloxopen/protoc-gen-preprocess/example/proto/*.proto /go/src/github.com/infobloxopen/protoc-gen-preprocess/example/proto/
RUN chmod +x /usr/bin/*
WORKDIR /go/src
