package main

import (
	"fmt"
	"telnet/internal/config"
	"telnet/internal/telnet"
)

func main() {
	config := config.New()
	fmt.Println(config)

	telnet := telnet.New(config)
	telnet.Run()

}
