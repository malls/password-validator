# Password Validator

## Usage

Password Validator is a command line utility that compares a list of passwords against a list of common, weak passwords. It also provides additional validation to ensure passwords are between 8 and 64 characters, and don't use non-ASCII characters.


`cat input_passwords.txt | ./password_validator weak_password_list.txt`



The output looks like this:

```
mom -> Error: Too Short
password1 -> Error: Too Common
*** -> Error: Invalid Charaters

```

## Installation







## Development

### Running Uncompiled
The main go entry point file is `src/main.go` and you can use this file much like the above example, piping input to it like this: `cat input.txt | go run src/main.go bad-passwords.txt`.

### Building

`go build src/main.go ./password-validator`

### Tests

