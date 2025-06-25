package handler

import (
	"github.com/onexstack/onex/internal/nightwatch/biz"
	"github.com/onexstack/onex/internal/nightwatch/pkg/validation"
)

// Handler implements a gRPC service.
type Handler struct {
	biz biz.IBiz
	val *validation.Validator
}

// NewHandler creates a new instance of Handler.
func NewHandler(biz biz.IBiz, val *validation.Validator) *Handler {
	return &Handler{biz: biz, val: val}
}
