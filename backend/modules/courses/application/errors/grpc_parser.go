package errs

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

var invalidAttribute *InvalidAttributeError
var useCaseError *UseCaseError
var notFoundError *NotFoundError
var permissionDeniedError *PermissionDeniedError

func NewGrpcError(err error) error {
	if errors.As(err, &invalidAttribute) {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if errors.As(err, &useCaseError) {
		parts := strings.Split(err.Error(), ":")
		if len(parts) > 1 {
			return status.Error(codes.Internal, fmt.Sprintf("%s.", parts[0]))
		}
		return status.Error(codes.Internal, err.Error())
	}

	if errors.As(err, &notFoundError) {
		return status.Error(codes.NotFound, err.Error())
	}

	if errors.As(err, &permissionDeniedError) {
		return status.Error(codes.PermissionDenied, err.Error())
	}

	return status.Error(codes.Internal, "Internal server error.")
}
