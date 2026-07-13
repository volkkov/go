# Technical Deep Dive

This document outlines the internal architecture and utility functions available in this repository.

## Concurrency Architecture
The project includes a custom worker pool pattern located in the `internal/worker` package.
* It manages a fixed number of worker goroutines for concurrent task processing.
* Tasks are enqueued via the `Submit` method, which safely blocks if the queue reaches capacity.
* A `Shutdown` method is provided to wait for running tasks to complete cleanly.

## Utility Helpers (root package main)

### String Utilities (`stringkit.go`)
Provides basic text manipulation functions:
* `Capitalize`: Converts the first letter of a string to uppercase.
* `SlugifyLite`: Generates URL-friendly slugs from strings.
* `TruncateWithEllipsis`: Shortens strings to a maximum length.

### Validation (`emailcheck.go` & `validator.go`)
The repository currently contains overlapping approaches to email validation:
* `IsEmailFormatValid`: Utilizes strict regular expressions for validation.
* `IsValidEmail`: Uses a simplified check ensuring the presence of an "@" and "." symbol.

### Application Constants (`constants.go`)
Global application configurations are stored here, defining the application version, maximum retries, and timeout limits.