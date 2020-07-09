# gitprompt-go

## Usage

This started out as a copy of gitprompt-rs (https://github.com/9ary/gitprompt-rs), but written in **go** instead of **rust**.

By default it will produce a string like this (depending on the directory you are in) :

    master↑0↓0|+0~3-0x0•0
    
 Explanation (copied from https://github.com/9ary/gitprompt-rs):
 
- Branch info:
  - `master`: name of the current branch, `:HEAD` in detached head mode
  - `↑`: number of commits ahead of remote
  - `↓`: number of commits behind remote
- Work area:
  - `+`: untracked (new) files
  - `~`: modified files
  - `-`: deleted files
  - `x`: merge conflicts
- `•`: staged changes

What I have added to 9ary:s code is a json config file, where you can change the prompt to match your situation. For example you can change the symbols used, reorder the information displayed, and change colors etc. The config file should be placed in `~/.config/gitprompt-go/config.json`.

A sample config file:
 ```
 {
   "format": "$(BRANCH)$(AHEAD)$(BEHIND)$(SEPARATOR)$(UNTRACKED)$(MODIFIED)$(DELETED)$(UNMERGED)$(STAGED)",
   "includeZeroValues" : false,
   "separator": "|",
   "promptPrefix": "(",
   "promptSuffix": ")",
   "branch" :  {
     "prefix": "$(ESC)[34m",
     "suffix": "$(ESC)[90m"
   },
   "ahead": {
     "prefix": "↑",
     "suffix": ""
   },
   "behind": {
     "prefix": "↓",
     "suffix": ""
   },
   "untracked": {
     "prefix": "+",
     "suffix": ""
   },
   "modified": {
     "prefix": "~",
     "suffix": ""
   },
   "deleted": {
     "prefix": "-",
     "suffix": ""
   },
   "unmerged": {
     "prefix": "x",
     "suffix": ""
   },
   "staged": {
     "prefix": "•",
     "suffix": ""
   }
 }
```

Use `$(ESC)` instead of `\x1b` or `\033`, it will be replaced by the code later.

After following Cris Titus Tech:s video (https://www.youtube.com/watch?v=iaXQdyHRL8M) and adding `$(gitprompt-go)` to your shell prompt, and creating the config file `config.json` above in the `~/.config/gitprompt-go` folder ,your prompt should look something like this:

![gitprompt-go](gitprompt.png)

## Known issues

* **gitprompt-rs** prompt work with ZSH, **gitprompt-go** probably doesn't. Since I don't use ZSH, I didn't bother to copy that code.
