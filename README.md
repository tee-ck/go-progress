# go-progress

## A simple progress bar for Golang CLI Application

# Installation

```shell
> go get -u github.com/tee-ck/go-progress
```

# Usage

```go
package main

import (
	"fmt"
	"github.com/tee-ck/go-progress"
	"math/rand"
	"time"
)

func main() {
	p := progress.New(1000*progress.MiB, progress.Config{
		Length: 30,
	})

	for !p.IsDone() {
		random := progress.DataUnit(rand.Uint64())

		p.Increase(2 * progress.MiB + random % (100 * progress.KiB))
		eta := p.Eta(p.AverageSpeed()).Sub(time.Now())
		fmt.Printf("\r[%s] %s / %s | %6.3f %%  | %s | %s | %s",
			p.Bar(),
			p.NowValue(),
			p.MaxValue(),
			p.Progress(),
			p.CurrentSpeed(),
			p.AverageSpeed(),
			eta,
		)

		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()
}
```

# Output
![]()

# License

The project is licensed under MIT license.