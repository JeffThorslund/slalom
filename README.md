# Go Kayak Race Analysis

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Description

This project provides a set of tools for analyzing kayak race data. It includes functionality for parsing race data, validating the data, and performing various calculations and analyses.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install this project, you will need to have Go installed on your machine. Once Go is installed, you can run this from the root directory with `go run .`

## Usage

This project includes several Go files that can be run individually:

- [`functions.go`](functions.go): Contains various utility functions used throughout the project.
- [`parsing.go`](parsing.go): Provides functionality for parsing race data from CSV files.
- [`slalom.go`](slalom.go): Contains the main logic for analyzing slalom race data.
- [`validation.go`](validation.go): Provides functions for validating the parsed data.

The project also includes a `.github/workflows/go.yml` file for GitHub Actions, which can be used to automatically build and test the project on every push or pull request.

## Contributing

Contributions to this project are welcome. Please ensure that any changes you make pass the existing tests and, if possible, add new tests for new functionality.

## License

This project is licensed under the [MIT License](LICENSE).