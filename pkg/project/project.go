package project

var (
	description = "A service that knows everything about your installation"
	gitSHA      = "n/a"
	name        = "athena"
	source      = "https://github.com/giantswarm/athena"
	version     = "1.13.1"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
