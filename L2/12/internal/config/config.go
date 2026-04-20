package config

import (
	"errors"
	"log"

	"github.com/spf13/pflag"
	flag "github.com/spf13/pflag"
)

var (
	errorNegativeArgs = errors.New("grep: context argument must be non-negative: Invalid argument")
)

type Config struct {
	Pattern  string
	FileName string

	AfterContext  int
	BeforeContext int
	Context       int
	Count         bool
	IgnoreCase    bool
	Invert        bool
	Fixed         bool
	LineNum       bool
}

func New() *Config {
	cfg := Config{}

	flag.IntVarP(&cfg.AfterContext, "after-context", "A", 0, "Print num lines of trailing context after each match")
	flag.IntVarP(&cfg.BeforeContext, "before-context", "B", 0, "Print num lines of leading context before each match")
	flag.IntVarP(&cfg.Context, "context", "C", 0, "Print num lines of leading and trailing context surrounding each match")
	flag.BoolVarP(&cfg.Count, "count", "c", false, "Only a count of selected lines is written to standard output")
	flag.BoolVarP(&cfg.IgnoreCase, "ignore-case", "i", false, "Perform case insensitive matching")
	flag.BoolVarP(&cfg.Invert, "invert-match", "v", false, "Selected lines are those not matching any of the specified patterns")
	flag.BoolVarP(&cfg.Fixed, "fixed-strings", "F", false, "Interpret pattern as a set of fixed strings")
	flag.BoolVarP(&cfg.LineNum, "line-number", "n", false, "Each output line is preceded by its relative line number in the file")

	pflag.Parse()

	if flag.NArg() > 0 {
		cfg.Pattern = flag.Arg(0)
		cfg.FileName = flag.Arg(1)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func (c *Config) Validate() error {
	if c.AfterContext < 0 || c.BeforeContext < 0 {
		return errorNegativeArgs
	}
	if c.Context < 0 {
		return errorNegativeArgs
	}

	return nil
}
