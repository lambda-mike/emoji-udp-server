package config

type Config struct {
	N         int
	Separator string
	Raw       bool
}

func Create(isRaw bool, n int, sep string) (Config, error) {
	// TODO validate params, return error if needed
	// n needs to be >= 0
	return Config{n, sep, isRaw}, nil
}
