# protoc-gen-preprocess

### Purpose
We have variety of validators for protobuf, however sometimes we need to pre-process incoming data. This protobuf plugin helps to generate preprocessing methods for incoming messages. 

### Example
```
syntax = "proto3";

import "github.com/infobloxopen/protoc-gen-preprocess/options/preprocess.proto";

message Demo {
   string preprocessedField = 1 [(preprocess.field).string.trim_space = true ];
   string untouched = 2;
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
	m.PreprocessedField = strings.TrimSpace(m.PreprocessedField)
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

### Running Demo

```
# build demo
make demo

# run example application
go run example/main.go
```

This will launch a server with our demo service. To check functionality of demo application send following request:
```
curl -X POST -i http://localhost:8080/echo --data '{"preprocessedField": "     Those spaces will be trimmed    ","untouched": " Notice how those spaces will be left as is    "}'
```

## Supported Field Types
For now following list of fields supported:
* **String**:
    *  **trim_spaces** - Will trim leading and following spaces