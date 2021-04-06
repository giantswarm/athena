package main

import (
	"github.com/giantswarm/microerror"
)

var cannotLoadConfigError = &microerror.Error{
	Kind: "cannotLoadConfigError",
}

var cannotGenerateCodeError = &microerror.Error{
	Kind: "cannotGenerateCodeError",
}
