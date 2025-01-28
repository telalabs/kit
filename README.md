# QQKit

![QQTbanner](https://github.com/user-attachments/assets/d89bfb72-2011-4fbf-b898-988108214842)


A dual-language package (Go/Rust) for building and managing LLM function calling tools and toolkits. Built specifically for QQ.

## Features

- Abstract Tool interface for implementing LLM-compatible functions
- Default tool implementation with required metadata
- Toolkit management for grouping related tools
- Functional options pattern for configuration (Go)
- Builder pattern for configuration (Rust)
- JSON schema support for function parameters and returns

## Installation

### Go
```bash
go get github.com/QQTechnologies/QQ/go
```

### Rust
```toml
[dependencies]
toolkit = { git = "https://github.com/QQTechnologies/kit", subdirectory = "rust" }
```

## Usage

Both implementations provide similar functionality with language-specific idioms:

- Go uses the functional options pattern for configuration
- Rust uses the builder pattern for configuration
- Both support async execution of tools
- Both use JSON schemas for parameter and return type definitions

For examples, see the language-specific directories:
- Go examples in [go/examples](go/examples)
- Rust examples in [rust/examples](rust/examples)
