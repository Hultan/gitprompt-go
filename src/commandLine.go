package main

import (
	"errors"
	"os"
)

type CommandLineParser struct {
	Help    bool
	Verbose bool
}

func NewCommandLineParser() *CommandLineParser {
	return new(CommandLineParser)
}

func (p *CommandLineParser) Parse() error {
	args := os.Args[1:]

	for _, value := range args {
		switch value {
		case "-h":
			p.Help = true
		case "-v":
			p.Verbose = true
		default:
			return errors.New("Invalid parameter : " + value)
		}
	}
	return nil
}
