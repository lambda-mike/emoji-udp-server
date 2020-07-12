package config

import (
	"errors"
	"strconv"
)

type Config struct {
	N         int
	Separator string
	Raw       bool
}

func Create(isRaw bool, n int, sep string) (Config, error) {
	if n < 0 {
		msg := "n must be positive, got: " + strconv.Itoa(n)
		return Config{}, errors.New(msg)
	}
	return Config{n, sep, isRaw}, nil
}

func ParsePort(port string) (int, error) {
	// TODO parse port
	return 0, errors.New("Error parsing port: " + port)
}
