package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	N         int
	Separator string
	Raw       bool
}

const EMOJI_PORT = "EMOJI_PORT"

func Create(isRaw bool, n int, sep string) (Config, error) {
	if n < 0 {
		msg := "n must be positive, got: " + strconv.Itoa(n)
		return Config{}, errors.New(msg)
	}
	return Config{n, sep, isRaw}, nil
}

func ReadPortFromEnv() string {
	return os.Getenv(EMOJI_PORT)
}

func ParsePort(str string) (int, error) {
	port, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	lo := 49152
	hi := 65535
	if port < lo || port > hi {
		msg := fmt.Sprintf("Port should be between %d-%d, got:%d", lo, hi, port)
		return 0, errors.New(msg)
	}
	return port, nil
}

func ParseCmdLineFlags() (Config, error) {
	// -h param is handled by the flag lib automatically
	isRaw := flag.Bool("r", false, "disable the translation from keyword to emoji")
	n := flag.Int("n", 1, "cmd input number multiplier")
	sep := flag.String("s", "", "emojis separator")
	flag.Parse()
	return Create(*isRaw, *n, *sep)
}
