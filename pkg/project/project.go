package project

var (
	description = "TODO"
	gitSHA      = "n/a"
	name        = "athena"
	source      = "https://github.com/giantswarm/athena"
	version     = "0.1.0-dev"
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
