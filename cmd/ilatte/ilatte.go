package main

import (
	"github.com/stny/latte/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
