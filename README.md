# Coding Katas Go 

![CI](https://github.com/eliflores/coding-katas-go/workflows/CI/badge.svg)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-%23FE5196?logo=conventionalcommits&logoColor=white)](https://conventionalcommits.org)

[Go](https://go.dev/) solutions for Katas from
* [LeetCode](https://leetcode.com/) 🧡
* And more... 🌈

## How to work with this repository

### On your local host 

1. Clone the repository
2. Install the version of [go](https://go.dev/) required in [go.mod](go.mod)
  * _Hint_: You may use [gvm](https://github.com/moovweb/gvm) to manage different versions of go.
3. Run: `go mod tidy`
4. Run: `go mod vendor`
5. Run the tests: `go test -v ./...`
6. Run the linting: `golangci-lint run`
   * Checkout the version of `golangci-lint` to use in the [CI workflow](.github/workflows/ci.yml).
7. Write the tests, write the code, lint the code, and have fun!
8. Open a Pull Request, wait for all the checks to pass, and merge! :tada:

✨ **Note** ✨: A [devcontainer.json](.devcontainer/devcontainer.json) file is available, so you could also 
work on a [Visual Studio Code Dev Container](https://code.visualstudio.com/docs/devcontainers/containers) without the need to install `go` or `golangci-lint` on your local host.

### On GitHub Codespaces 

You can also try out opening the repository in a [GitHub Codespace](https://github.com/features/codespaces) 🌈
