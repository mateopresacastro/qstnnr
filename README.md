# qstnnr - A Go Quiz Application

A command-line quiz application for a take home assignment built with `Go`, featuring a `gRPC` API server and CLI client. Users can take quizzes, get feedback, and compare their performance against other participants.

## Features

- `gRPC` API Server
- Command-line interface for taking the quiz
- In-Memory Storage
- Performance comparison with other participants

## Technical Stack

- Backend: `Go`
- `API`: `gRPC` with `Protocol Buffers`ß
- Storage: In-memory
- CLI Framework: `Cobra`
- Error Handling: Domain-specific error types with stack traces
- Testing: Unit and integration tests

## Getting Started

### Prerequisites

- `Go`
- Optional: `Protocol Buffers` compiler (`protoc`) and `Make`.

### Installation

Clone the repository:

```bash
git clone https://github.com/mateopresacastro/qstnnr.git && cd qstnnr
```

Build the application:

```bash
make build
```

This will create two binaries in the `bin` directory:

`server`: The `gRPC` server

`qstnnr`: The CLI client

## Running the Application

To run the quiz run `qstnnr server start` and then `qstnnr take`:

```bash
➜ bin/qstnnr server start
Server started on port 4000 (PID: 12152)
```

```bash
➜ bin/qstnnr take
Question 1 of 10
Use the arrow keys to navigate: ↓ ↑ → ←
What is the purpose of the blank identifier (_) in Go?
  ➜ To discard an unwanted value
    To declare a private variable
    To create an anonymous function
    To mark a variable as nullable
```

The CLI has two main commands: `server` and `take`.

```bash
➜ bin/qstnnr help
A CLI application to check you Go knowledge.

Usage:
  qstnnr [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  server      Manage the qstnnr server
  take        Take the quiz

Flags:
  -h, --help   help for qstnnr

Use "qstnnr [command] --help" for more information about a command.
```

### `server` command

To see the help for the server command run:

```bash
➜ bin/qstnnr server help
Commands to start, stop, restart and check status of the qstnnr server

Usage:
  qstnnr server [command]

Available Commands:
  restart     Restart the qstnnr server
  start       Start the qstnnr server
  status      Check the status of the qstnnr server
  stop        Stop the qstnnr server

Flags:
  -h, --help   help for server

Use "qstnnr server [command] --help" for more information about a command.
```

You can set `PORT` in your `env` vars before starting the server. By default it uses `5974`.

```bash
➜ bin/qstnnr server start
Server started on port 5974 (PID: 95853)
```

```bash
➜ export PORT=4000
➜ bin/qstnnr server start
Server started on port 4000 (PID: 96592)
```

## `take` command

The `take` command starts the quiz. At the end you can see your results.

```bash
➜ bin/qstnnr take
Question 1 of 10
Use the arrow keys to navigate: ↓ ↑
What is the zero value for a pointer in Go?
  ➜ 0
    undefined
    void
    nil
```

## Project Structure

```bash
├── cmd/ # Application entrypoints
│ ├── cli/ # CLI implementation
│ └── server/ # Server implementation
├── pkg/
│ ├── api/ # gRPC protocol definitions
│ ├── qerr/ # Error handling
│ ├── qservice/ # Business logic
│ ├── server/ # gRPC server implementation
│ └── store/ # Data storage
├── Makefile # Build and development commands
├── questions.go          # Quiz content and initial data
└── run.go               # Main application setup and server initialization
```

## Features in Detail

### Quiz Taking Flow

- User requests questions through CLI
- Server returns questions with multiple choice options
- User submits answers
- Server evaluates answers and returns:
  - Number of correct answers
  - Comparison with other participants
  - Correct solutions

### Error Handling

- Domain-specific error types with stack traces
- Graceful error propagation through layers
- Proper `gRPC` status code mapping

### Configuration

Environment variables:

- `PORT`: Server port (default: `5974`)
- `LOG_LEVEL`: Logging level (`DEBUG`, `INFO`, `WARN`, `ERROR`)

## Development

### Building

```bash
make build # Build all binaries
```

```bash
make build-server # Build server only
```

```bash
make build-cli # Build CLI only
```

### Testing

- Comprehensive test coverage
- Integration tests for `gRPC` server

```bash
make test # Run all tests with coverage and race detection
```

### Generating Protocol Buffers

```bash
make proto # Generate Go code from .proto files
```
