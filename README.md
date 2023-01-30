# ZHS Logger

Logger implements logrus with custom default formatter.

Example:
```go
package main

import (
	"github.com/zhs/zhs-logger"
)

func main() {
    log := logger.New()

    log.Info("Log message")
}
```

Output:
```
2023-01-30 20:55:19 [INFO ]: test
```