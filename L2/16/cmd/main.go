package main

import (
	"wget/internal/config"
	"wget/internal/downloader"
)

func main() {
	cfg := config.New()

	downloader := downloader.New(cfg)
	downloader.Start()
}
