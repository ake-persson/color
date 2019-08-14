[![GoDoc](https://godoc.org/github.com/mickep76/color?status.svg)](https://godoc.org/github.com/mickep76/color)
[![Go Report Card](https://goreportcard.com/badge/github.com/mickep76/color)](https://goreportcard.com/report/github.com/mickep76/color)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mickep76/color/blob/master/LICENSE)

# color

Terminal color package for Go


## Example

```go
package main
  
import (
        "fmt"

        "github.com/mickep76/color"
)

func main() {
        fmt.Printf("%s%s%s%s\n", color.Red, color.Bold, "Hello World!", color.Reset)

        for i := 1; i < 255; i++ {
                fmt.Printf("%s#", color.Fg256(uint8(i)))
        }
        fmt.Printf("%s\n", color.Reset)

        for i := 1; i < 255; i++ {
                fmt.Printf("%s#", color.Bg256(uint8(i)))
        }
        fmt.Printf("%s\n", color.Reset)
}
```
