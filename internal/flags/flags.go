package flags

type Flags struct {
	Address        string
	AllowedOrigins []string
}

func New() *Flags {
	f := &Flags{}

	return f
}
