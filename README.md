# go-binary-version-manager (gobvm)

Version manager for go libraries, currently golangci-lint and mockery support.

[![🔨Build And Test](https://github.com/Abdulsametileri/go-binary-version-manager/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/Abdulsametileri/go-binary-version-manager/actions/workflows/test.yml)

# Motivation

The projects I worked on used different versions of the `golangci-lint` and `mockery` libraries. When I use their
commands, inconsistent, strange errors occur.
I am aiming to solve this problem by writing a basic binary manager that uses symlink under the hood.
This project infra is easy to extend for other libraries too :)  

# Demo

[![asciicast](https://asciinema.org/a/663612.svg)](https://asciinema.org/a/663612)

# Installation

## Brew

`brew install abdulsametileri/tap/gbvm`

## Golang

`go install github.com/Abdulsametileri/go-binary-version-manager@latest`

After installing go command, the binary artifact name is `go-binary-version-manager`, 
You can rename it like
`mv $GOPATH/bin/go-binary-version-manager $GOPATH/bin/gbvm` 
or
`mv $HOME/go/bin/go-binary-version-manager $HOME/go/bin/gbvm`

# Commands

| Command                          | Explanation                                          | Example                              |
|----------------------------------|------------------------------------------------------|--------------------------------------|
| gobvm listall $LIBRARY           | it lists all installed versions of the given library | `gobvm listall mockery`              |
| gobvm enable $LIBRARY@$VERSION   | it enables given version of the library              | `gobvm enable golangci-lint@v1.55.1` |
| gobvm install $LIBRARY@VERSION   | it installs given version of the library             | `gobvm install mockery@v2.20.0`      |
| gobvm uninstall $LIBRARY@VERSION | it uninstalls given version of the library           | `gobvm uninstall mockery@v2.20.0`    |


# TODO
- [ ] e2e test like [vx](https://github.com/Abdulsametileri/vX/blob/main/e2e-test.sh)