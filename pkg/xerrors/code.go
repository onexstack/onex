// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/superproj/miniblog.

package xerrors

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// Code is the type that represents an error code.
// It can map to HTTP and gRPC codes.
// In order to properly work with custom codes or code overrides:
// use the RegisterCode function after creating your new Code instance.
type Code struct {
	grpc   codes.Code
	http   int
	reason string
}

var (
	OK Code = Code{codes.OK, http.StatusOK, ""}
	// Unknown is the default error code.
	Unknown Code = Code{codes.Unknown, http.StatusInternalServerError, "Unknown"}
	// Internal is used when an internal error occurs.
	Internal Code = Code{codes.Internal, http.StatusInternalServerError, "InternalError"}
	// InvalidArgument is used when the client sends invalid arguments.
	InvalidArgument Code = Code{codes.InvalidArgument, http.StatusBadRequest, "InvalidArgument"}
	// NotFound is used when the requested resource is not found.
	NotFound Code = Code{codes.NotFound, http.StatusNotFound, "NotFound"}
	// AlreadyExists is used when the resource already exists.
	AlreadyExists Code = Code{codes.AlreadyExists, http.StatusConflict, "AlreadyExist"}
	// Unauthorized is used when the client is not authenticated.
	Unauthenticated Code = Code{codes.Unauthenticated, http.StatusUnauthorized, "Unauthenticated"}
	// PermissionDenied is used when the client is not authorized to perform the requested operation.
	PermissionDenied Code = Code{codes.PermissionDenied, http.StatusForbidden, "PermissionDenied"}
)

func NewCode(grpcCode codes.Code, httpCode int, reason string) Code {
	return Code{
		grpc:   grpcCode,
		http:   httpCode,
		reason: reason,
	}
}

func (c Code) R(reason string) Code {
	c.reason = reason
	return c
}

// String returns the string representation of the code.
func (c Code) String() string {
	return c.reason
}

// HTTP returns the HTTP code that is mapped to the code.
func (c Code) HTTP() int {
	return c.http
}

// GRPC returns the gPRC code that is mapped to the code.
func (c Code) GRPC() codes.Code {
	return c.grpc
}

// MarshalJSON implements the json.Marshaler interface and defines how a Code
// should be marshaled to JSON. By default, it marshals to a string representation defined by String function.
func (c Code) MarshalJSON() ([]byte, error) {
	s := c.String()
	return []byte("\"" + s + "\""), nil
}
