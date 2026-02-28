package config

import (
	"flag"
)

type Config struct {
	URL   string
	Depht int
}

func New() *Config {
	url := flag.String("url", "", "url")
	depht := flag.Int("d", 1, "depht")

	flag.Parse()

	return &Config{
		URL:   *url,
		Depht: *depht,
	}
}
