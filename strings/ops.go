package main

import (
  "fmt"
  "unicode/utf8"
)

func main() {
  const word = "Hello, 世界"
  fmt.Println("Using the UTF8 decoder:")
  for i := 0; i < len(word); {
    r, size := utf8.DecodeRuneInString(word[i:])
    fmt.Printf("%d\t%c\t%d\n", i, r, size)
    i += size
  }
  fmt.Println("Using built in range (preferred):")
  for i, r := range word {
    fmt.Printf("%d\t%q\t%d\n", i, r, r)
  }
}
