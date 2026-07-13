# Go Playground

A collection of Go utility functions and a basic application entry point. This repository serves as a playground for various Go patterns, including concurrent workers, string manipulation, and mathematical operations.

## Prerequisites
* [Go](https://golang.org/doc/install) installed on your local machine.

## Running the Project
The primary entry point for the application is `hello.go`. It contains a simple main function that prints a greeting.

To run the application, navigate to the root directory in your terminal and execute:
\`\`\`bash
go run hello.go
\`\`\`

## Project Structure
This repository contains several helper files in the root `main` package, plus an internal worker package used for concurrency:
* `stringkit.go`: Helpers for string formatting.
* `mathutils.go`: Functions for calculating sums and averages.
* `internal/worker/pool.go`: A bounded goroutine pool for concurrent task execution.

For a detailed technical breakdown, please refer to `docs/deep-dive.md`.