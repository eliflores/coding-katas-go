# Coding Katas Go 

![CI](https://github.com/eliflores/coding-katas-python/workflows/CI/badge.svg)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

[Go](https://go.dev/) solutions for Katas from
* [LeetCode](https://leetcode.com/) 🧡
* And more... 🌈

## How to work with this repository

1. Clone the repository
2. Install the version of go required in [go.mod](go.mod)
3. Run: `go mod tidy`
4. Run: `go mod vendor`
5. Run the tests: `go test -v ./...`
6. Run the linting: `golangci-lint run`
   * Checkout the version of `golangci-lint` to use in the [CI workflow](.github/workflows/ci.yml).
7. Write the tests, write the code, lint the code, and have fun!
8. Open a Pull Request, wait for all the checks to pass, and merge! :tada:
