# stats_cli

CLI tool for analyzing CSV data

### Installation
To install the stats_cli package as an executable use:
    `go install github.com/Meowcenary/stats_cli@latest`
After installing commands can be run with:
    `stats_cli <command> [options]`

To clone the stats_cli package use:
    `git clone https://github.com/Meowcenarystats_benchmark.git`
Once cloned commands can be run from the root directory:
    `go run main.go <command> [options]`

### Running the tests
To run all the tests use `go test ./...` from the root directory of the project.
To run a single package's tests use `go test ./<package name>` E.g
`go test ./analysis`

To run the benchmark use `go test -bench=.`. Pass the flag `-count` to run the
benchmark repeatedly. For counts over 700 the flag `-timeout` will need to be
passed or the benchmark will quit before finishing all runs. Here is an example
of a full benchmarking command for 1000 runs:
```
go test -bench=. -count=1000 -timeout 20m
```

### Using the CLI
Currently the only command available is `summary` which reads a csv file of
decimal and integer numbers from the path provided by the argument `--file`.
File is the only requird argument, the rest are optional.

All available arguments for `stats_cli summary`:
- `--file` - CSV file to read data from
- `--columns` - Columns to include from CSV for summary output
- `--stats` - Stats to include for summary output

Example usage:
```
stats_cli summary --file testdata.csv --columns age,income --stats max,min,mean
```
Or from the project root with
```
go run main.go summary --file testdata.csv --columns age,income --stats max,min,mean
```

### Benchmarking Results

#### Golang

Details for running the Golang benchmark can be found above.  The average
benchmark time over 100 runs was 0.09062145 ns.

#### Python/Pandas

The Python script was benchmarked using the library `timeit`. The average
benchmark time over 100 runs was 10.194 ns.

#### R

The R script was benchmarked using a simple system time comparison. The average
benchmark time over 100 runs was 73757629 ns (0.073757629 s).
