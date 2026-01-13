package config

import (
	"flag"
)

type Config struct {
	Fields    string
	Delimiter string
	Separated bool
}

func New() *Config {
	fields := flag.String("f", "", "fields")
	delimiter := flag.String("d", "\t", "delimiter")
	sep := flag.Bool("s", false, "separated")

	flag.Parse()

	return &Config{
		Fields:    *fields,
		Delimiter: *delimiter,
		Separated: *sep,
	}
}
