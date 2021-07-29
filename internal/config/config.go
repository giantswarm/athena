package config

type Flags struct {
	Server       Server
	Identity     Identity
	Kubernetes   Kubernetes
	Capabilities Capabilities
}

func New() *Flags {
	f := &Flags{}

	return f
}
