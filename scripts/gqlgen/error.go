package main

import (
	"errors"

	"github.com/giantswarm/microerror"
)

var cannotLoadConfigError = &microerror.Error{
	Kind: "cannotLoadConfigError",
}

// IsCannotLoadConfig asserts cannotLoadConfigError.
func IsCannotLoadConfig(err error) bool {
	return errors.Is(err, cannotLoadConfigError)
}

var cannotGenerateCodeError = &microerror.Error{
	Kind: "cannotGenerateCodeError",
}

// IsCannotGenerateCode asserts cannotGenerateCodeError.
func IsCannotGenerateCode(err error) bool {
	return errors.Is(err, cannotGenerateCodeError)
}
