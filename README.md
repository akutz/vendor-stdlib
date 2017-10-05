# Vendor-stdlib
Vendor-stdlib illustrates how it is possible to build Go programs with
small changes to the Go runtime by vendoring stdlib.

This example adds a new function "StringsEqFold" to the stdlib package
"sort". The new function sorts string slices using a case-insensitive
comparison. For example:

```go
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
```

The above program accepts one or more arguments on the command line
and prints them in increasing order using a case-insensitive comparison:

```bash
$ go run main.go hi there Andrew!
Andrew!
hi
there
```
