package api

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// The canonical errors from the EigenDA gRPC API endpoints.
//
// Notes:
// - We start with a small (but sufficient) subset of grpc's error codes,
//   and expand when there is an important failure case to separate out. See:
//   https://grpc.io/docs/guides/status-codes/
// - Make sure that internally propagated errors are eventually wrapped in one of the
//   user-facing errors defined here, since grpc otherwise returns an UNKNOWN error code,
//   which is harder to debug and understand for users.
// - See https://github.com/googleapis/googleapis/blob/ba8ea80f25d19bde8501cd51f314391f8d39bde8/google/rpc/code.proto
//   for the mapping of grpc error codes to HTTP status codes.

func newErrorGRPC(code codes.Code, msg string) error {
	return status.Error(code, msg)
}

// HTTP Mapping: 400 Bad Request
func NewErrorInvalidArg(msg string) error {
	return newErrorGRPC(codes.InvalidArgument, msg)
}

// HTTP Mapping: 404 Not Found
func NewErrorNotFound(msg string) error {
	return newErrorGRPC(codes.NotFound, msg)
}

// HTTP Mapping: 429 Too Many Requests
func NewErrorResourceExhausted(msg string) error {
	return newErrorGRPC(codes.ResourceExhausted, msg)
}

// HTTP Mapping: 500 Internal Server Error
func NewErrorInternal(msg string) error {
	return newErrorGRPC(codes.Internal, msg)
}

// HTTP Mapping: 500 Internal Server Error
func NewErrorUnknown(msg string) error {
	return newErrorGRPC(codes.Unknown, msg)
}

// HTTP Mapping: 501 Not Implemented
func NewErrorUnimplemented() error {
	return newErrorGRPC(codes.Unimplemented, "not implemented")
}

// HTTP Mapping: 504 Gateway Timeout
func NewErrorDeadlineExceeded(msg string) error {
	return newErrorGRPC(codes.DeadlineExceeded, msg)
}

func NewErrorCanceled(msg string) error {
	return newErrorGRPC(codes.Canceled, msg)
}

// ErrorFailover is returned by the disperser-client and eigenda-client to signify
// that eigenda is temporarily unavailable, and suggest to the caller
// (most likely some rollup batcher via the eigenda-proxy) to failover
// to ethda for some amount of time.
// See https://github.com/ethereum-optimism/specs/issues/434 for more details.
//
// Given that both clients already return grpc errors, we could potentially use
// a grpc UNAVAILABLE error instead, but we don't because:
//  1. UNAVAILABLE is typically used to tell the client to retry the request, not failover
//  2. the grpc framework itself also returns UNAVAILABLE errors in some cases, see:
//     https://github.com/grpc/grpc-go/blob/192ee33f6fc0f07070eeaaa1d34e41746740e64c/codes/codes.go#L184.
//     We could differentiate from those generated by the grpc framework by using error details, like
//     https://github.com/grpc/grpc-go/tree/master/examples/features/error_details, but that would complicate things
//     and it feels much simpler to just use a custom error type for this specific purpose.
//
// 3 reasons for returning api.ErrorFailover:
//  1. Failed to put the blob in the disperser's queue (disperser is down)
//  2. Timed out before getting confirmed onchain (batcher is down)
//  3. Insufficient signatures (eigenda network is down)
//
// One can check if an error is an ErrorFailover by using errors.Is:
//
//	failoverErr := NewErrorFailover(someOtherErr)
//	if !errors.Is(wrappedFailoverErr, &ErrorFailover{}) {
//	  // do something...
//	}
type ErrorFailover struct {
	Err error
}

// NewErrorFailover creates a new ErrorFailover with the given underlying error.
// See ErrorFailover for more details.
func NewErrorFailover(err error) *ErrorFailover {
	return &ErrorFailover{
		Err: err,
	}
}

func (e *ErrorFailover) Error() string {
	if e.Err == nil {
		return "Failover"
	}
	return fmt.Sprintf("Failover: %s", e.Err.Error())
}

func (e *ErrorFailover) Unwrap() error {
	return e.Err
}

// Is only checks the type of the error, not the underlying error.
// This is because we want to be able to check that an error is an ErrorFailover,
// even when wrapped. This can now be done with errors.Is.
//
//	baseErr := fmt.Errorf("some error")
//	failoverErr := NewErrorFailover(baseErr)
//	wrappedFailoverErr := fmt.Errorf("some extra context: %w", failoverErr)
//
//	if !errors.Is(wrappedFailoverErr, &ErrorFailover{}) {
//	  // do something...
//	}
func (e *ErrorFailover) Is(target error) bool {
	_, ok := target.(*ErrorFailover)
	return ok
}
