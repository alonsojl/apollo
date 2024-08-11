package internal

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	Alphabetic   = validation.Match(regexp.MustCompile("^[a-zA-ZáéíóúñÁÉÍÓÚÑ ]+$")).Error("must contain Spanish letters")
	Alphanumeric = validation.Match(regexp.MustCompile("^[a-zA-Z0-9áéíóúñÁÉÍÓÚÑ _.,]+$")).Error("must contain Spanish letters and digits only")
)
