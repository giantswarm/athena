package analytics

import (
	"errors"

	"github.com/giantswarm/microerror"
)

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return errors.Is(err, invalidConfigError)
}

var validationError = &microerror.Error{
	Kind: "validationError",
}

// IsValidation asserts validationError.
func IsValidation(err error) bool {
	return errors.Is(err, validationError)
}
