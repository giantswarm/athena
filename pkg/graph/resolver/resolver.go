package resolver

import (
	"github.com/giantswarm/microerror"
	"go.uber.org/zap"

	"github.com/giantswarm/athena/pkg/certificate"
)

//go:generate rm -rf generated
//go:generate go run ../../../scripts/gqlgen

type ResolverConfig struct {
	Log *zap.SugaredLogger

	InstallationProvider   string
	InstallationCodename   string
	InstallationK8sApiUrl  string
	InstallationK8sAuthUrl string
	InstallationK8sCaCert  string
}

type Resolver struct {
	log *zap.SugaredLogger

	installationProvider   string
	installationCodename   string
	installationK8sApiUrl  string
	installationK8sAuthUrl string
	installationK8sCaCert  string
}

func NewResolver(config ResolverConfig) (*Resolver, error) {
	if config.Log == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Log must not be empty", config)
	}

	r := &Resolver{
		log:                    config.Log,
		installationProvider:   config.InstallationProvider,
		installationCodename:   config.InstallationCodename,
		installationK8sApiUrl:  config.InstallationK8sApiUrl,
		installationK8sAuthUrl: config.InstallationK8sAuthUrl,
		installationK8sCaCert:  certificate.Parse(config.InstallationK8sCaCert),
	}

	return r, nil
}
