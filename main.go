package main

import (
	"fmt"
	"os"

	"github.com/jharrels/perses/invalidator"
)

func main() {
	inv := invalidator.FromArgs(os.Args[1:])
	if err := inv.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %v\n", err)
		os.Exit(1)
	}
}
