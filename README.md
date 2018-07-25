## Balance

This package primarily provides a function for testing whether an input string has a balanced set of braces ('{' and '}'). The data structure used to perform the balance test is also exported.

This package can be imported into your own code with something similar to the code sample below:
```golang
package main

import (
	"fmt"

	"github.com/subtlepseudonym/balance"
)

func main() {
	valid := "{hello}{world}"
	invalid := "{}}{}"

	validOutput := balance.Balance(valid) // returns -1
	invalidOutput := balance.Balance(invalid) // returns 2

	fmt.Printf("valid balance: %d", validOutput)
	fmt.Printf("invalid balance: %d", invalidOutput)
}
```

You can test this code using the go toolchain's built-in test function:
```bash
go test balance.go balance_test.go
```

You can benchmark this code by adding a few additional flags:
```bash
go test balance.go balance_test.go --bench . --benchmem
```
