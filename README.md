# primeapp-testing
Testing Primeapp on Go

Command to check test coverage.

go test -coverprofile=coverage.out && go tool cover -html=coverage.out

Tips for test prompt()

1) fmt.Print write string to os.Stdout 

2) we can temporary change / replace os.Stdout with os.Pipe()
