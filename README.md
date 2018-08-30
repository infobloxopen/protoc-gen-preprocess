# protoc-gen-preprocess

### Purpose
We have variety of validators for protobuf, however sometimes we need to pre-process incoming data. This protobuf plugin helps to generate preprocessing methods for incoming messages. 

### Example
```
syntax = "proto3";

import "github.com/infobloxopen/protoc-gen-preprocess/options/preprocess.proto";

message Demo {
   string s = 1 [(preprocess.field).string.trim_space = true ];
}
```
will generate:
```
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/infobloxopen/protoc-gen-preprocess/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *Demo) Preprocess() error {
	m.S = strings.TrimSpace(m.S)
	return nil
}
```

## Usage
It is best to use this plugin as a middleware as shown in example application:

```
package main

import 	(
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

```
# get repo
go get -d github.com/infobloxopen/protoc-gen-preprocess

# build
make install

```