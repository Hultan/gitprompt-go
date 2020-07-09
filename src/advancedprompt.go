package main

//# Reset
//Color_Off='\033[0m'       # Text Reset
//
//# Regular Colors
//Black='\033[0;30m'        # Black
//Red='\033[0;31m'          # Red
//Green='\033[0;32m'        # Green
//Yellow='\033[0;33m'       # Yellow
//Blue='\033[0;34m'         # Blue
//Purple='\033[0;35m'       # Purple
//Cyan='\033[0;36m'         # Cyan
//White='\033[0;37m'        # White
//
//# Bold
//BBlack='\033[1;30m'       # Black
//BRed='\033[1;31m'         # Red
//BGreen='\033[1;32m'       # Green
//BYellow='\033[1;33m'      # Yellow
//BBlue='\033[1;34m'        # Blue
//BPurple='\033[1;35m'      # Purple
//BCyan='\033[1;36m'        # Cyan
//BWhite='\033[1;37m'       # White
//
//# Underline
//UBlack='\033[4;30m'       # Black
//URed='\033[4;31m'         # Red
//UGreen='\033[4;32m'       # Green
//UYellow='\033[4;33m'      # Yellow
//UBlue='\033[4;34m'        # Blue
//UPurple='\033[4;35m'      # Purple
//UCyan='\033[4;36m'        # Cyan
//UWhite='\033[4;37m'       # White
//
//# Background
//On_Black='\033[40m'       # Black
//On_Red='\033[41m'         # Red
//On_Green='\033[42m'       # Green
//On_Yellow='\033[43m'      # Yellow
//On_Blue='\033[44m'        # Blue
//On_Purple='\033[45m'      # Purple
//On_Cyan='\033[46m'        # Cyan
//On_White='\033[47m'       # White
//
//# High Intensity
//IBlack='\033[0;90m'       # Black
//IRed='\033[0;91m'         # Red
//IGreen='\033[0;92m'       # Green
//IYellow='\033[0;93m'      # Yellow
//IBlue='\033[0;94m'        # Blue
//IPurple='\033[0;95m'      # Purple
//ICyan='\033[0;96m'        # Cyan
//IWhite='\033[0;97m'       # White
//
//# Bold High Intensity
//BIBlack='\033[1;90m'      # Black
//BIRed='\033[1;91m'        # Red
//BIGreen='\033[1;92m'      # Green
//BIYellow='\033[1;93m'     # Yellow
//BIBlue='\033[1;94m'       # Blue
//BIPurple='\033[1;95m'     # Purple
//BICyan='\033[1;96m'       # Cyan
//BIWhite='\033[1;97m'      # White
//
//# High Intensity backgrounds
//On_IBlack='\033[0;100m'   # Black
//On_IRed='\033[0;101m'     # Red
//On_IGreen='\033[0;102m'   # Green
//On_IYellow='\033[0;103m'  # Yellow
//On_IBlue='\033[0;104m'    # Blue
//On_IPurple='\033[0;105m'  # Purple
//On_ICyan='\033[0;106m'    # Cyan
//On_IWhite='\033[0;107m'   # White
//
//And then use them like this in your script:
//
//#    .---------- constant part!
//#    vvvv vvvv-- the code from above
//RED='\033[0;31m'
//NC='\033[0m' # No Color
//printf "I ${RED}love${NC} Stack Overflow\n"

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
	parts := strings.Split(a.Config.Format, "$(SEPARATOR)")

	result := a.getPromptPart(parts[0])
	if len(parts)>1 {
		var part string
		for i := 1; i < len(parts); i++ {
			part = a.getPromptPart(parts[i])
			if len(result)>0 && len(part)>0 {
				result = strings.Join([]string{result, part}, a.Config.Separator)
			}
		}
	}

	return a.escape(a.Config.PromptPrefix + result + a.Config.PromptSuffix)
}

func (a *AdvancedPrompt) escape(text string) string {
	return strings.ReplaceAll(text,"$(ESC)","\x1b")
}

func (a *AdvancedPrompt) getPromptPart(part string) string {
	result := part

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