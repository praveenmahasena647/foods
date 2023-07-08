package main

import (
	"fmt"
	"os"

	"github.com/praveenmahasena647/foodApi/api"
)

func main() {
	var err error = api.RunServer()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
