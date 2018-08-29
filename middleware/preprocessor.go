package grpc_preprocessor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type preprocessor interface {
	Preprocess() error
}

// UnaryServerInterceptor returns a new unary server interceptor that preprocesses incoming messages.
//
// Invalid messages will be rejected with `InvalidArgument` before reaching any userspace handlers.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(preprocessor); ok {
			if err := v.Preprocess(); err != nil {
				return nil, grpc.Errorf(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	}
}
