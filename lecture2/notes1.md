## Lecture 2

The os package in Go provides a way to interact with the operating system.
functionalities provided by the os package:

1. Environment Variables:
    `os.Getenv(name string) string`: Retrieves the value of the environment variable with the given name.

2. Command-Line Arguments:
    `os.Args`: A slice of strings that represents the command-line arguments.

3. Working with Files and Directories:
    `os.Create(name string) (*os.File, error)`: Creates or truncates a file with the specified name.

    `os.Open(name string) (*os.File, error)`: Opens a file for reading.

    `os.OpenFile(name string, flag int, perm FileMode) (*os.File, error)`: Opens a file using a specified flag and permissions.

4. Standard Input/Output/Error:
    `os.Stdin`, `os.Stdout`, `os.Stderr`: Variables representing standard input, output, and error.

5. Removing Files or Directories. etc.




> Note: We should import the hello package using the module path[check the go.mod file]
It threw an error, when I did:<br>
`lecture2/hello`<br>
But worked fine, when I wrote:<br>
`github.com/Praveen005/matt-holiday-go-lectures/tree/main/lecture2/hello`



- To run the test file:
    - All the tests in PWD:
        `go test ./...`
    
    - Run all tests in the hello directory:
        `go test ./hello/...`

    - Run specific test file:
        `go test ./hello/my_test.go`

    - Run specific test function:
        `go test ./hello/my_test.go:TestMyFunction`

- Use the `-v` flag for verbose output to see more details about the test execution.<br>
    `go test -v ./...`

- Use the `-cover` flag to generate a coverage report showing which parts of your code are covered by tests.<br>
    `go test -cover ./...`


> **Do you know:** Go has a separate [GoMoney](https://pkg.go.dev/github.com/Rhymond/go-money#section-readme) package, which help us with precise Money operations such as rounding, splitting and allocating.
Monetary values should not be stored as floats due to small rounding differences.