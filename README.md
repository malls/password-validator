# Password Validator

## Usage

Password Validator is a command line utility that compares a list of passwords against a list of common, weak passwords. It also provides additional validation to ensure passwords are between 8 and 64 characters, and don't use non-ASCII characters.


`cat input.txt | password_validator invalidlist.txt`


The output looks like this:

```
mom -> Error: Too Short
password1 -> Error: Too Common
*** -> Error: Invalid Charaters

```

Developed with Go `go1.10.2`.

## Installation

Copy the `password-validator` binary to somewhere on your `$PATH`, for example `cp password-validator ~/usr/local/bin`. (This location can vary). If it doesn't run, try building (see below) and try this again, as you may need a version built for your specific system. 

## Development

### Running Uncompiled
The main go entry point file is `src/main.go` and you can use this file much like the above example, piping input to it like this: `cat input.txt | go run src/main.go bad-passwords.txt`. `src/main.go` is the main entry point for the program. Supporting packages are found in folders within `src/` with names matching the package name.


### Tests
You can run the unit tests with `go test -v ./...`. Everything must pass before submitting a pull request. If you add a new package, also add `packagename_test.go` within the package folder.

### Building

`go build src/main.go` will build compile the code and replace `password-validator`.



