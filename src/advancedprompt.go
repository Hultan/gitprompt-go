package main

import (
	"strconv"
	"strings"
)

type AdvancedPrompt struct {
	Git    *Git
	Config *Config
}

func NewAdvancedPrompt(git *Git, config *Config) *AdvancedPrompt {
	prompt := new(AdvancedPrompt)
	prompt.Git = git
	prompt.Config = config
	return prompt
}

func (a *AdvancedPrompt) GetPrompt() string {
	var result string

	result = a.Config.Format

	result = strings.ReplaceAll(result, "$(SEPARATOR)",a.Config.Separator)
	branch := a.getBranch(a.Git, a.Config)
	result = strings.ReplaceAll(result, "$(BRANCH)",branch)
	if a.Config.IncludeZeroValues || a.Git.Ahead>0 {
		ahead := a.getAhead(a.Git, a.Config)
		result = strings.ReplaceAll(result, "$(AHEAD)",ahead)
	} else {
		result = strings.ReplaceAll(result, "$(AHEAD)","")
	}
	if a.Config.IncludeZeroValues || a.Git.Behind>0 {
		behind := a.getBehind(a.Git, a.Config)
		result = strings.ReplaceAll(result, "$(BEHIND)", behind)
	} else {
		result = strings.ReplaceAll(result, "$(BEHIND)", "")
	}
	if a.Config.IncludeZeroValues || a.Git.Untracked>0 {
		untracked := a.getUntracked(a.Git, a.Config)
		result = strings.ReplaceAll(result, "$(UNTRACKED)", untracked)
	} else {
		result = strings.ReplaceAll(result, "$(UNTRACKED)", "")
	}
	if a.Config.IncludeZeroValues || a.Git.Modified>0 {
		modified := a.getModifed(a.Git, a.Config)
		result = strings.ReplaceAll(result, "$(MODIFIED)", modified)
	} else {
		result = strings.ReplaceAll(result, "$(MODIFIED)", "")
	}
	if a.Config.IncludeZeroValues || a.Git.Deleted>0 {
		deleted := a.getDeleted(a.Git, a.Config)
		result = strings.ReplaceAll(result, "$(DELETED)", deleted)
	} else {
		result = strings.ReplaceAll(result, "$(DELETED)", "")
	}
	if a.Config.IncludeZeroValues || a.Git.Unmerged>0 {
		unmerged := a.getUnmerged(a.Git, a.Config)
		result = strings.ReplaceAll(result, "$(UNMERGED)", unmerged)
	} else {
		result = strings.ReplaceAll(result, "$(UNMERGED)", "")
	}
	if a.Config.IncludeZeroValues || a.Git.Staged>0 {
		staged := a.getStaged(a.Git, a.Config)
		result = strings.ReplaceAll(result, "$(STAGED)", staged)
	} else {
		result = strings.ReplaceAll(result, "$(STAGED)", "")
	}

	return result
}

func (a *AdvancedPrompt) getBranch(git *Git, config *Config) string {
	return config.Branch.Prefix + git.Branch + config.Branch.Suffix
}

func (a *AdvancedPrompt) getAhead(git *Git, config *Config) string {
	return config.Ahead.Prefix + strconv.Itoa(git.Ahead) + config.Ahead.Suffix
}

func (a *AdvancedPrompt) getBehind(git *Git, config *Config) string {
	return config.Behind.Prefix + strconv.Itoa(git.Behind) + config.Behind.Suffix
}

func (a *AdvancedPrompt) getUntracked(git *Git, config *Config) string {
	return config.Untracked.Prefix + strconv.Itoa(git.Untracked) + config.Untracked.Suffix
}

func (a *AdvancedPrompt) getModifed(git *Git, config *Config) string {
	return config.Modified.Prefix + strconv.Itoa(git.Modified) + config.Modified.Suffix
}
func (a *AdvancedPrompt) getDeleted(git *Git, config *Config) string {
	return config.Deleted.Prefix + strconv.Itoa(git.Deleted) + config.Deleted.Suffix
}
func (a *AdvancedPrompt) getUnmerged(git *Git, config *Config) string {
	return config.Unmerged.Prefix + strconv.Itoa(git.Unmerged) + config.Unmerged.Suffix
}
func (a *AdvancedPrompt) getStaged(git *Git, config *Config) string {
	return config.Staged.Prefix + strconv.Itoa(git.Staged) + config.Staged.Suffix
}