# RePet - CLI Snippet Manager

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://github.com/knqyf263/pet/blob/master/LICENSE)


# Motivation

`RePet` is a fork of the [pet](https://github.com/knqyf263/pet) cli tool.

`RePet` adds the functionality of repeating the previous command called, and the ability to put a command into the clip board after executing.

# Additional features
## Repeating last command
The `--repeat` or `-r` flag to the `exec` command which repeats the last executed command using the same parameter values as the one used previously. 
```bash
$ repet exec
> echo "hello"
hello
$ repet exec -r
> echo "hello"
hello
``` 

## Excluded snippets
`RePet` stores the latest executed command in the file system. Some commands might end up containing information (like passwords) that should not be stored in the file system. Such snippets can be marked as `Excluded`. When a snippet is marked as `Excluded` then they won't be stored, and and the `repeat` flag will not execute them.
```bash
$ repet exec
> echo "hello"
hello
$ repet exec
> excluded snippet
$ repet exec -r
> echo "hello"
hello
``` 

## Copy command to clipboard
The `--copy` or `-c` flag for the `exec` command can be used to copy the executed command into clipboard when it's executed. 
```bash
$ repet exec -c
> echo "hello"
hello
$ echo "hello"
```