package external

import (
	"io"
	"os"
	"os/exec"
)

func Execute(name string, args []string, stdin io.Reader, stdout io.Writer) (*exec.Cmd, error) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return cmd, nil
}
