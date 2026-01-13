package main

import (
	"cut/internal/config"
	"cut/internal/cut"
	"fmt"
	"os"
)

func main() {
	config := config.New()
	fmt.Println(config)

	srs := os.Stdin

	cutter := cut.New(config, srs)
	err := cutter.DoCut()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
