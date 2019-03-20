FROM golang:1.10.0 AS builder

LABEL stage=server-intermediate

WORKDIR /go/src/github.com/infobloxopen/protoc-gen-preprocess
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/usr/bin/protoc-gen-preprocess main.go

FROM infoblox/atlas-gentool:latest AS runner

COPY --from=builder /out/usr/bin/protoc-gen-preprocess /usr/bin/protoc-gen-preprocess
COPY --from=builder /go/src/github.com/infobloxopen/protoc-gen-preprocess/options/*.proto /go/src/github.com/infobloxopen/protoc-gen-preprocess/options
WORKDIR /go/src
