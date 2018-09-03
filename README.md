# protoc-gen-preprocess

## Purpose

We have variety of validators for protobuf, however sometimes we need to pre-process incoming data. This protobuf plugin helps to generate preprocessing methods for incoming messages.

### Example

```proto
syntax = "proto3";

import "github.com/infobloxopen/protoc-gen-preprocess/options/preprocess.proto";

message Demo {
   string preprocessedField = 1 [(preprocess.field).string.trim_space = true ];
   repeated string preprocessedRepeatedField = 2 [(preprocess.field).string.trim_space = true ];
   string untouched = 3;
}

```

will generate:

```go
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-preprocess/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *Demo) Preprocess() error {

    m.PreprocessedField = strings.TrimSpace(m.PreprocessedField)

    for i, s := range m.PreprocessedRepeatedField {
        m.PreprocessedRepeatedField[i] = strings.TrimSpace(s)
    }
    return nil
}

```

## Usage

It is best to use this plugin as a middleware as shown in example application:

```go
package main

import (
...
    grpc_preprocessor "github.com/infobloxopen/protoc-gen-preprocess/middleware"
...
)

func runService() {
...
// Middleware chain.
    interceptors := []grpc.UnaryServerInterceptor{
...
        grpc_preprocessor.UnaryServerInterceptor(), // preprocessing middleware
...
    }
    server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))
...
}

```

### Installation

```sh
# get repo
go get -d github.com/infobloxopen/protoc-gen-preprocess

# build
make install

```

### Running Demo

```sh
# build demo
make demo

# run example application
go run example/main.go
```

This will launch a server with our demo service. To check functionality of demo application send following request:

```sh
curl -X POST -i http://localhost:8080/echo --data '{"preprocessedField": "     Those spaces will be trimmed    ","untouched": " Notice how those spaces will be left as is    "}'
```

## Supported Field Types

For now following list of fields supported:

* **String**:
  * **trim_space** - Will trim leading and following spaces