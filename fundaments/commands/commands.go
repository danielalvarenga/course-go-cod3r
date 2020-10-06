package main

import "fmt"

func main() {
	fmt.Printf("Another %s program!\n", "go")
}

/*
Terminal commands:

Help about get
go help get

Import packages
go get -u github.com/go-sql-driver/mysql
That will be imported from github to GOPATH/src/github.com/go-sql-driver/mysql

Offilne go doc
godoc -http=:6060

Environments
go env

Detect problems in the program
go vet commands.go

Compile e create executable file
go build commands.go

Compile and execute
go run commands.go
*/
