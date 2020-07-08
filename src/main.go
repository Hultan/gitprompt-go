package main

import "fmt"

func main() {
	// Create the git prompt object
	prompt:=NewGitPrompt()
	// Print the git prompt
	fmt.Println(prompt.Prompt)
}
