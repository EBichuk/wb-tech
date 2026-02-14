package builtins

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

type Builtins struct {
	builtins map[string]struct{}
}

func New() *Builtins {
	return &Builtins{
		builtins: map[string]struct{}{
			"cd":   {},
			"pwd":  {},
			"echo": {},
			"kill": {},
			"ps":   {},
		},
	}
}

func (b *Builtins) IsBuiltins(command string) bool {
	if _, ok := b.builtins[command]; ok {
		return ok
	}
	return false
}

func (b *Builtins) RunBuiltin(command string, out io.Writer, args []string) error {
	switch command {
	case "cd":
		return b.changeDirectory(args)
	case "pwd":
		return b.printWorkingDirectory(out)
	case "echo":
		return b.echo(out, args[0])
	case "kill":
		return b.kill(args)
	case "ps":
		return b.ps(out, args)
	}

	return nil
}

func (b *Builtins) changeDirectory(args []string) error {
	var path string

	switch len(args) {
	case 0:
		path = os.Getenv("HOME")
	case 1:
		path = args[0]
	default:
		return fmt.Errorf("cd: too many arguments %v", args)
	}

	if err := os.Chdir(path); err != nil {
		return fmt.Errorf("cd %w", err)
	}

	return nil
}

func (b *Builtins) printWorkingDirectory(out io.Writer) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("pwd: %w", err)
	}
	_, err = fmt.Fprint(out, dir)
	return err
}

func (b *Builtins) echo(out io.Writer, args string) error {
	_, err := fmt.Fprint(out, args)
	return err
}

func (b *Builtins) kill(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("kill: usage: kill pid")
	}

	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("kill: %s arguments must be process or job IDs", args[0])
	}

	err = syscall.Kill(pid, syscall.SIGTERM)
	if err != nil {
		return fmt.Errorf("kill: operation not permitted")
	}

	return nil
}

func (b *Builtins) ps(out io.Writer, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("ps: too many arguments")
	}

	c := exec.Command("ps", args...)
	c.Stdin = os.Stdin
	c.Stdout = out
	c.Stderr = os.Stderr
	return c.Run()
}
