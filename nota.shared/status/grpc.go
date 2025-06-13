package error

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var StatusInternal = status.Error(codes.Internal, "something went wrong, try again later")
