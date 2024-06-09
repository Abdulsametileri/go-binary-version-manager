# go-binary-version-manager (gobvm)

Version manager for go libraries, currently golangci-lint and mockery support.

# Motivation

The projects I worked on used different versions of the `golangci-lint` and `mockery` libraries. When I use their
commands, inconsistent, strange errors occur.
I am aiming to solve this problem by writing a basic binary manager that uses symlink under the hood.
This project infra is easy to extend for other libraries too :)  

# Commands

| Command                          | Explanation                                          | Example                              |
|----------------------------------|------------------------------------------------------|--------------------------------------|
| gobvm listall $LIBRARY           | it lists all installed versions of the given library | `gobvm listall mockery`              |
| gobvm enable $LIBRARY@$VERSION   | it enables given version of the library              | `gobvm enable golangci-lint@v1.55.1` |
| gobvm install $LIBRARY@VERSION   | it installs given version of the library             | `gobvm install mockery@v2.20.0`      |
| gobvm uninstall $LIBRARY@VERSION | it uninstalls given version of the library           | `gobvm uninstall mockery@v2.20.0`    |

# TODO
- [ ] Listall, Uninstall cmd implementation (golangci-lint)
- [ ] Listall, Install, Enable, Uninstall implementation (mockery)
- [ ] Brew package
- [ ] Makefile, lint
- [ ] PIPELINE
- [ ] ASCINEMA & Doc
- [ ] e2e test like vx