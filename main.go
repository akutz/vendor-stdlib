package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: at least one token is required")
		os.Exit(1)
	}
	toks := os.Args[1:]
	sort.StringsEqFold(toks)
	for _, t := range toks {
		fmt.Println(t)
	}
}
