package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"shell/internal/builtins"
	"shell/internal/external"
	"shell/internal/parser"
	"sync"
	"syscall"
)

type Shell struct {
	parser                 parser.Parser
	builtins               builtins.Builtins
	sigCh                  chan os.Signal
	mu                     sync.Mutex
	externalRunningCommand []*exec.Cmd
}

func New() *Shell {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	return &Shell{
		parser:   *parser.New(),
		builtins: *builtins.New(),
		sigCh:    sigCh,
		mu:       sync.Mutex{},
	}
}

func (s *Shell) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	go s.SignalCancel()

	for scanner.Scan() {

		command := scanner.Text()
		if command == "" {
			continue
		}

		cmd := s.parser.Parse(command)

		err := s.Execute(cmd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}

}

func (s *Shell) SignalCancel() {
	for range s.sigCh {
		s.mu.Lock()
		for _, cmd := range s.externalRunningCommand {
			if cmd != nil && cmd.Process != nil {
				err := cmd.Process.Signal(syscall.SIGINT)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v", err)
				}
			}
		}
		s.mu.Unlock()
	}
}

func (s *Shell) Execute(cmds []parser.Command) error {
	var err error

	if len(cmds) < 2 {
		err = s.executeSingleCmd(cmds[0])
	} else {
		err = s.executePipeline(cmds)
	}
	return err
}

func (s *Shell) executeSingleCmd(cmd parser.Command) error {
	if s.builtins.IsBuiltins(cmd.Name) {
		err := s.builtins.RunBuiltin(cmd.Name, os.Stdout, cmd.Args)
		if err != nil {
			return err
		}

	} else {
		execCmd, err := external.Execute(cmd.Name, cmd.Args, os.Stdin, os.Stdout)
		if err != nil {
			return err
		}

		if err := execCmd.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Shell) executePipeline(pipeline []parser.Command) error {
	readerFile := make([]*os.File, len(pipeline))
	writerFile := make([]*os.File, len(pipeline))

	for i := range len(pipeline) - 1 {
		r, w, err := os.Pipe()
		if err != nil {
			closeResources(r, w)
			return err
		}

		readerFile[i+1], writerFile[i] = r, w
	}

	readerFile[0] = os.Stdin
	writerFile[len(pipeline)-1] = os.Stdout

	for i, cmd := range pipeline {
		if s.builtins.IsBuiltins(cmd.Name) {
			err := s.builtins.RunBuiltin(cmd.Name, writerFile[i], cmd.Args)
			if err != nil {
				closeResources(readerFile...)
				closeResources(writerFile...)
				return err
			}
		} else {
			execCmd, err := external.Execute(cmd.Name, cmd.Args, readerFile[i], writerFile[i])
			if err != nil {
				closeResources(readerFile...)
				closeResources(writerFile...)
				return err
			}

			s.mu.Lock()
			s.externalRunningCommand = append(s.externalRunningCommand, execCmd)
			s.mu.Unlock()
		}
	}

	closeResources(readerFile...)
	closeResources(writerFile...)

	var lastError error
	for i := range s.externalRunningCommand {
		if err := s.externalRunningCommand[i].Wait(); err != nil {
			lastError = err
		}
	}

	s.mu.Lock()
	s.externalRunningCommand = []*exec.Cmd{}
	s.mu.Unlock()

	return lastError
}

func closeResources(resources ...*os.File) {
	for _, resource := range resources {
		if resource != nil && resource != os.Stdin && resource != os.Stdout {
			_ = resource.Close()
		}
	}
}
