[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/sure/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/sure/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/sure)](https://pkg.go.dev/github.com/yyle88/sure)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/sure/master.svg)](https://coveralls.io/github/yyle88/sure?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/sure.svg)](https://github.com/yyle88/sure/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/sure)](https://goreportcard.com/report/github.com/yyle88/sure)

# sure: Add Assertions and Crash Handling to Existing Go Code

`sure` enhances your existing Go code by adding assertions and crash handling. It automatically asserts conditions and crashes when errors occur, allowing you to improve error handling in legacy code without needing to manually add repetitive checks.

## CHINESE README

[中文说明](README.zh.md)

## CREATION_IDEAS

[CREATION_IDEAS](internal/docs/CREATION_IDEAS.en.md)

## Packages Overview

### `sure_cls_gen`: **Generates Go Classes with Assertions**

Generates Go classes from predefined objects, embedding assertion logic to prevent common errors.

### `sure_pkg_gen`: **Generates Go Packages with Error Handling**

Extracts functions from existing code and generates Go packages, integrating assertion and crash handling.

### `cls_stub_gen`: **Generates Package-Level Function Wrappers**

Creates package-level functions that wrap methods of a singleton struct, simplifying access usage.

## Usage

### Examples:

- [Generating Classes with `sure_cls_gen`](internal/examples/example_sure_cls_gen)
- [Generating Packages with `sure_pkg_gen`](internal/examples/example_sure_pkg_gen)
- [Generating Singleton with `cls_stub_gen`](internal/examples/example_cls_stub_gen)

---

## License

`sure` is open-source and released under the MIT License. See the LICENSE file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `sure`!** 🎉

Give me stars. Thank you!!!

## See stars
[![see stars](https://starchart.cc/yyle88/sure.svg?variant=adaptive)](https://starchart.cc/yyle88/sure)
