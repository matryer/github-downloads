package main

import (
	"fmt"
	"os"

	ghdownloads "github.com/matryer/github-downloads"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("provide repo in form: username/reponame")
		os.Exit(1)
	}
	n, err := ghdownloads.Count(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Print(n)
}
