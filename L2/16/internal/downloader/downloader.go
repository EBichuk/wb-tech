package downloader

import "wget/internal/config"

type Downloader struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Downloader {
	return &Downloader{
		cfg: cfg,
	}
}

func (d *Downloader) Start() {

}
