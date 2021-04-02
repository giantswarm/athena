package config

type Flags struct {
	Server     Server
	Identity   Identity
	Kubernetes Kubernetes
}

func New() *Flags {
	f := &Flags{}

	return f
}
