package internal

import (
	"fmt"
	"strconv"
)

type GitPrompt struct {
	Prompt string
}

func NewGitPrompt() *GitPrompt {
	prompt := new(GitPrompt)
	prompt.getPrompt()
	return prompt
}

func (g *GitPrompt) getPrompt() {
	// Parse the command line
	parser:= NewCommandLineParser()
	err := parser.Parse()
	if err != nil {
		g.Prompt = err.Error() + "\n\n" + g.getUsage()
		return
	}

	// If the user have asked for help, give it to them
	if parser.Help {
		g.Prompt = g.getUsage()
		return
	}

	// Get the git status
	git:= NewGitStatus()

	// If it is not a git repository, just leave
	if !git.IsGit {
		g.Prompt = ""
		return
	}

	// The user have asked for a verbose git status
	if parser.Verbose {
		g.Prompt = g.getVerbosePrompt(git)
		return
	}

	config := NewConfig()
	if config.ConfigExists() {
		config.Load()
		advanced := NewAdvancedPrompt(git, config)
		g.Prompt = advanced.GetPrompt()
	} else {
		g.Prompt = g.getDefaultPrompt(git)
	}
}

func (g *GitPrompt) getVerbosePrompt(git *Git) string {
	var result string

	result += fmt.Sprintf("Branch    : %s (%d ahead, %d behind)\n", git.Branch, git.Ahead, git.Behind)
	result += fmt.Sprintf("Staged    : %d\n", git.Staged)
	result += fmt.Sprintf("Modified  : %d\n", git.Modified)
	result += fmt.Sprintf("Deleted   : %d\n", git.Deleted)
	result += fmt.Sprintf("Unmerged  : %d\n", git.Unmerged)
	result += fmt.Sprintf("Untracked : %d\n", git.Untracked)

	return result
}

func (g *GitPrompt) getUsage() string {
	return "Usage : gitprompt-go [-h] [-d] [-v]"
}

func (g *GitPrompt) getDefaultPrompt(git *Git) string {
	// Create and return the normal git prompt
	var result string
	if git.Branch!="" {
		// Branch
		result = git.Branch
	} else {
		// Detached head
		result = ":HEAD"
	}
	if git.Ahead>0 {
		result += "↑" + strconv.Itoa(git.Ahead)
	}
	if git.Behind>0 {
		result += "↓" + strconv.Itoa(git.Behind)
	}
	if git.Untracked + git.Modified + git.Deleted + git.Unmerged + git.Staged>0 {
		result += "|"
	}
	if git.Untracked>0 {
		result += "+" + strconv.Itoa(git.Untracked)
	}
	if git.Modified>0 {
		result += "~" + strconv.Itoa(git.Modified)
	}
	if git.Deleted>0 {
		result += "-" + strconv.Itoa(git.Deleted)
	}
	if git.Unmerged>0 {
		result += "x" + strconv.Itoa(git.Unmerged)
	}
	if git.Staged>0 {
		result += "•" + strconv.Itoa(git.Staged)
	}

	return result
}