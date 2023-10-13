# stats_cli

CLI tool for analyzing CSV data

### Installation
To install the stats_cli package as an executable use:
    `go install github.com/Meowcenary/stats_cli@latest`

To clone the stats_cli package use:
    `git clone https://github.com/Meowcenarystats_benchmark.git`

### Running the tests
To run all the tests use `go test ./...` from the root directory of the project.
To run a single package's tests use `go test ./<package name>` E.g
`go test ./analysis`

### Using the CLI
From the root directory of the proejct use: `go run main.go <command> [options]`
Currently the only command available is `summary` which reads a csv file of
decimal and integer numbers from the path provided by the argument `--file`.
File is the only requird argument, the rest are optional.

All available arguments for `stats_cli summary`:
- `--file` - CSV file to read data from
- `--columns` - Columns to include from CSV for summary output
- `--stats` - Stats to include for summary output
