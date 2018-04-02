# Clock

Example of usage:

```go
package main

import (
	"fmt"
	"github.com/xrash/clock"
	"time"
)

func main() {
	c := clock.NewClock()
	c.Start()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			fmt.Println("running...", c.Now())
		}
	}()

	time.Sleep(time.Second * 5)
	c.Stop()

	fmt.Println("stopped...", c.Now())
}
```
