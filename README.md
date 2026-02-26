[![Go](https://github.com/Gumarapus-com/go_pkg_tutor/actions/workflows/go.yml/badge.svg)](https://github.com/Gumarapus-com/go_pkg_tutor/actions/workflows/go.yml)

# Go-Pkg-Tutor

This repository purpose just for Golang package development demo.

> Here is the article link: [LINK TURORIAL](https://gumarapus.com/blog/golang-package-development)

## Usage

Add module

```sh
go get github.com/Gumarapus-com/go_pkg_tutor
```

`main.go`

```go
package main

import "github.com/Gumarapus-com/go_pkg_tutor"

func main() {
    greeting.Greet("My Name")
}
```

Run `go mod tidy`

```sh
# Add the required modules
go mod tidy

# Run the code
go run main.go
```

Result

```sh
Greeting "My name"
```
