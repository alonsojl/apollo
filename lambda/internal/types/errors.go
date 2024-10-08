package types

import (
	"apollo/pkg/errorx"
)

const (
	CodeInvalidArgument errorx.ErrorCode = iota
	CodeUnauthorized
	CodeNotFound
)

var (
	ErrInvalidParams = errorx.NewErrorf(CodeInvalidArgument, "invalid params")
	ErrInvalidBody   = errorx.NewErrorf(CodeInvalidArgument, "invalid body")
)
