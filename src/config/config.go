package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	N         uint
	Separator string
	Raw       bool
}

const EMOJI_PORT = "EMOJI_PORT"

func Create(isRaw bool, n uint, sep string) Config {
	return Config{n, sep, isRaw}
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
	if port != 0 && (port < lo || port > hi) {
		msg := fmt.Sprintf("Port should be 0 or between %d-%d, got:%d", lo, hi, port)
		return 0, errors.New(msg)
	}
	return port, nil
}

func ParseCmdLineFlags() Config {
	// -h param is handled by the flag lib automatically
	isRaw := flag.Bool("r", false, "disable the translation from keyword to emoji")
	n := flag.Uint("n", 1, "cmd input number multiplier")
	sep := flag.String("s", "", "emojis separator")
	flag.Parse()
	return Create(*isRaw, *n, *sep)
}
