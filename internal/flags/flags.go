package flags

type Flags struct {
	Server Server
}

func New() *Flags {
	f := &Flags{}

	return f
}
