package main

import (
	"fmt"
	"gitprompt-go/internal"
)

func main() {
	// Create the git prompt object
	prompt := internal.NewGitPrompt()

	// Print the git prompt
	fmt.Println(prompt.Prompt)
}
