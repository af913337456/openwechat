# Random

[![Go Reference](https://pkg.go.dev/badge/github.com/tamboto2000/random.svg)](https://pkg.go.dev/github.com/tamboto2000/random)

Random is a small library for generating random strings.

# Features

  - Generate hexadecimals string
  - Generate string with options (include uppercase, numbers, and symbols)

### Installation

```sh
$ go get github.com/tamboto2000/random
```

### Example
```go
package main

import (
	"fmt"

	"github.com/tamboto2000/random"
)

func main() {
	// generate random string
	fmt.Println(random.RandStr(20))

	// generate random string with option
	fmt.Println(random.RandStrWithOpt(20, random.Option{
		IncludeNumber:    true,
		IncludeUpperCase: true,
		// IncludeSymbols: true,
	}))

	// generate random hexadecimals string
	fmt.Println(random.RandHexStr(20))
}

```

License
----

MIT
