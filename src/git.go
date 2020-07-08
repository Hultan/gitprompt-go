package main

import (
	"os/exec"
	"strconv"
	"strings"
)

type Git struct {
	IsGit bool
	Error error

	Branch string
	Ahead  int
	Behind int

	Staged    int
	Modified  int
	Deleted   int
	Unmerged  int
	Untracked int
}

const (
	GitStatusCommand = "git status --porcelain=v2 -z --branch --untracked-files=all"
	NilCharacter     = "\x00"
)

func NewGitStatus() *Git {
	g := new(Git)
	g.getStatus()
	return g
}

func (g *Git) getStatus() {
	status, err := g.callGitStatus()
	if err != nil {
		g.Error = err
		return
	}
	g.parseGitStatus(status)
	g.IsGit = true
}

func (g *Git) callGitStatus() (string, error) {
	command := exec.Command("/bin/bash", "-c", GitStatusCommand)
	out, err := command.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (g *Git) parseGitStatus(status string) {
	items := strings.Split(status, NilCharacter)
	for _, value := range items {
		switch {
		case strings.HasPrefix(value, "#"):
			g.parseBranch(value)
		case strings.HasPrefix(value, "1"), strings.HasPrefix(value, "2"):
			g.parseFile(value)
		case strings.HasPrefix(value, "u"):
			g.Unmerged += 1
		case strings.HasPrefix(value, "?"):
			g.Untracked += 1
		default:
		}
	}
}

func (g *Git) parseBranch(branch string) {
	items := strings.Split(branch, " ")
	switch items[1] {
	case "branch.head":
		if items[2] != "(detached)" {
			g.Branch = items[2]
		}
	case "branch.ab":
		ahead, err := strconv.Atoi(items[2])
		if err != nil {
			g.Error = err
			return
		}
		g.Ahead = ahead

		behind, err := strconv.Atoi(items[3])
		if err != nil {
			g.Error = err
			return
		}
		g.Behind = behind
	}
}

func (g *Git) parseFile(file string) {
	items := strings.Split(file, " ")
	fileStatus := items[1]

	if fileStatus[0] != '.' {
		g.Staged += 1
	}

	switch {
	case fileStatus[1] == 'M':
		g.Modified += 1
	case fileStatus[1] == 'D':
		g.Modified += 1
	}
}
