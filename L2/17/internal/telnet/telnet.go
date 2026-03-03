package telnet

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"telnet/internal/config"
	"time"
)

const readTimeout = time.Second

type Telnet struct {
	cfg *config.Config
}

func New(config *config.Config) *Telnet {
	return &Telnet{
		cfg: config,
	}
}

func (t *Telnet) Run() {

	address := net.JoinHostPort(t.cfg.Host, t.cfg.Port)
	conn, err := net.DialTimeout("tcp", address, t.cfg.Timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "telnet: error to dial tcp connection: %v", err)
		return
	}
	defer conn.Close()

	wg := &sync.WaitGroup{}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// writer
	wg.Go(func() {
		defer cancel()
		writer := bufio.NewWriter(conn)
		reader := bufio.NewReader(os.Stdin)

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			line, err := reader.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Fprint(os.Stderr, "Connection slosed.\n")
				} else {
					fmt.Fprintf(os.Stderr, "telnet: error of reading from stdin: %v", err)
				}
				return
			}

			if _, err := writer.WriteString(line); err != nil {
				fmt.Fprintf(os.Stderr, "telnet: error to write: %v", err)
				return
			}
			if err := writer.Flush(); err != nil {
				fmt.Fprintf(os.Stderr, "telnet: error to flush: %v", err)
				return
			}
		}
	})

	wg.Go(func() {
		defer cancel()
		reader := bufio.NewReader(conn)
		buffer := make([]byte, 2048)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			if err := conn.SetReadDeadline(time.Now().Add(readTimeout)); err != nil {
				fmt.Fprintf(os.Stderr, "telnet: error to set read deadline: %v", err)
				return
			}

			n, err := reader.Read(buffer)
			if err != nil {
				var netErr net.Error
				if errors.As(err, &netErr) && netErr.Timeout() {
					continue
				}
				if errors.Is(err, io.EOF) {
					fmt.Fprint(os.Stdout, "Connection slosed.\n")
				} else {
					fmt.Fprintf(os.Stderr, "telnet: error to read from socket: %v", err)
				}
				return
			}

			if n > 0 {
				fmt.Fprint(os.Stdout, string(buffer[:n]))
			}
		}
	})

	wg.Go(func() {
		select {
		case <-ctx.Done():
			return
		case <-sigCh:
			cancel()
			return
		}
	})

	wg.Wait()
}
