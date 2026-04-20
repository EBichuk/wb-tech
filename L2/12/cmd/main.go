package main

import (
	"grep/internal/config"
	"grep/internal/grep"
	"os"
)

func main() {
	cnf := config.New()

	grep := grep.New(cnf, os.Stdout)
	grep.Run()
}
