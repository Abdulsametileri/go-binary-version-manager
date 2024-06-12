# go-binary-version-manager (gbvm)

Version manager for go libraries, currently golangci-lint and mockery support.

[![🔨Build And Test](https://github.com/Abdulsametileri/go-binary-version-manager/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/Abdulsametileri/go-binary-version-manager/actions/workflows/test.yml)

[![Go Report Card](https://goreportcard.com/badge/github.com/Abdulsametileri/go-binary-version-manager)](https://goreportcard.com/report/github.com/Abdulsametileri/go-binary-version-manager)

# Motivation

The projects I worked on used different versions of the `golangci-lint` and `mockery` libraries. When I use their
commands, inconsistent, strange errors occur.
I am aiming to solve this problem by writing a basic binary manager that uses symlink under the hood.
This project infra is easy to extend for other libraries too :)  

**Note**: Library binaries must be within ($GOPATH/go/bin) before using gbvm. You can delete and install it via gbvm.

# Demo

[![asciicast](https://asciinema.org/a/663612.svg)](https://asciinema.org/a/663612)

# Installation

### via Brew

`brew install abdulsametileri/tap/gbvm`

### via Golang

`go install github.com/Abdulsametileri/go-binary-version-manager@latest`

After installing go command, the binary artifact name is `go-binary-version-manager`, 
you can rename it like

```shell
mv $GOPATH/bin/go-binary-version-manager $GOPATH/bin/gbvm
``` 

or

```shell
mv $HOME/go/bin/go-binary-version-manager $HOME/go/bin/gbvm
```

# Behind the scenes

![behind-the-scenes.png](.github%2Fimages%2Fbehind-the-scenes.png)

# Commands

| Command                         | Explanation                                          | Example                             |
|---------------------------------|------------------------------------------------------|-------------------------------------|
| gbvm listall $LIBRARY           | it lists all installed versions of the given library | `gbvm listall mockery`              |
| gbvm enable $LIBRARY@$VERSION   | it enables given version of the library              | `gbvm enable golangci-lint@v1.55.1` |
| gbvm install $LIBRARY@VERSION   | it installs given version of the library             | `gbvm install mockery@v2.20.0`      |
| gbvm uninstall $LIBRARY@VERSION | it uninstalls given version of the library           | `gbvm uninstall mockery@v2.20.0`    |


# TODO
- [ ] e2e test like [vx](https://github.com/Abdulsametileri/vX/blob/main/e2e-test.sh)