> Thank you for your interest in contributing! 

## Prerequisites

Before getting started with development please make sure to:

- **1. Install Required Dependencies:**
    - `go` *v1.25.0* (duh)
    - `gofumpt` *v0.9.2* (formatter)
    - `golangci-lint` *v2.6.0* (linter) 

> [!NOTE]
> If you have [mise](https://mise.jdx.dev/) this can all be done by running `mise install` in project
> root

- **2. Run `go mod tidy`**

## Creating a command 
Hkup uses cobra-cli library from creating the CLI. Please review
[documentation](https://cobra.dev/)
for best practices.

Additional project-specific requirements include:
- putting command logic in the `internal/logic` package
- making sure commands always run and return errors when appropriate via `RunE`
- adding testing for new commands
    - there is the `internal/testutil` package that should use
- following the file name convention for commands with subcommands:
    Ex. a command `echo` with subcommands `hello` and `bye` should have files:
    - *echo.go* 
    - *echo-hello.go*
    - *echo-bye.go*
