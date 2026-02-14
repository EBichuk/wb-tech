package parser

import (
	"strings"
)

type Parser struct {
}

type Command struct {
	Name string
	Args []string
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input string) []Command {
	inputCmd := strings.Fields(input)

	resCmd := make([]Command, 0)

	pointerNameCmd := 0
	for i := 0; i < len(inputCmd); i++ {

		if inputCmd[i] == "|" {

			argsLst := make([]string, i-pointerNameCmd-1)
			copy(argsLst, inputCmd[pointerNameCmd+1:i])

			resCmd = append(resCmd, Command{
				Name: inputCmd[pointerNameCmd],
				Args: argsLst,
			})
			pointerNameCmd = i + 1
		}
	}

	argsLst := make([]string, len(inputCmd)-pointerNameCmd-1)
	copy(argsLst, inputCmd[pointerNameCmd+1:])

	resCmd = append(resCmd, Command{
		Name: inputCmd[pointerNameCmd],
		Args: argsLst,
	})

	return resCmd
}
