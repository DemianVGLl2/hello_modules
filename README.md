# Hello Modules

This is an example of Go Modules

## Installation

Execute the following command:

````bash
go mod init test_git
go get -u github.com/DemianVGLl2/hello_modules
````

## Usage

````go
package main

import (
	"fmt"

	"github.com/DemianVGLl2/hello_modules"
)

func main() {
	message := hello_modules.Hello2("Demián")
	fmt.Printf(message)

	fmt.Printf(hello_modules.RandomHello(), "Demián")
}

````