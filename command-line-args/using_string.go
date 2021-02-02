package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  // fmt.Println(strings.Join(os.Args[1:], " "))
  for idx, item := range os.Args {
    fmt.Println(idx, item)
  }
  fmt.Println(strings.Join(os.Args, " "))
}
