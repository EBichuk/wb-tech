package config

import (
	"flag"
	"time"
)

type Config struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func New() *Config {
	host := flag.String("h", "localhost", "host")
	port := flag.String("p", "2525", "port")
	timeout := flag.Duration("timeout", time.Duration(time.Second*10), "timeout - default 10 seconds")

	flag.Parse()

	return &Config{
		Host:    *host,
		Port:    *port,
		Timeout: *timeout,
	}
}
