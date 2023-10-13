## stats_cli

CLI tool for analyzing CSV data

### Installation
`git clone https://github.com/Meowcenarystats_benchmark.git`

### Running the tests
To run all the tests use `go test ./...` from the root directory of the project.
To run a single package's tests use `go test ./<package name>` E.g
`go test ./analysis`

### Using the CLI
From the root directory of the proejct use: `go run main.go <command> [options]`
Currently the only command available is `summary` which reads a csv file of
decimal and integer numbers from a hardcoded path and prints a simple
statistical analysis of the data.
