package main

import (
  "fmt"
  "strings"
)

func main () {
  const wordA = "a/b/c.go"
  const wordB = "c.d.go"
  const wordC = "abc"

  fmt.Println("Method A: ")
  fmt.Println(basename(wordA))
  fmt.Println(basename(wordB))
  fmt.Println(basename(wordC))

  fmt.Println("Method B: ")
  fmt.Println(basenameV2(wordA))
  fmt.Println(basenameV2(wordB))
  fmt.Println(basenameV2(wordC))
}

func basename(s string) string {
  // Discard last '/' and everything before .
  for i := len(s) - 1; i >=0; i-- {
    if s[i] == '/' {
      s = s[i + 1:]
      break
    }
  }
  // Preserve everything before last '.'
  for i := len(s) - 1; i >= 0; i -- {
    if s[i] == '.' {
      s = s[:i]
      break
    }
  }
  return s
}

func basenameV2(s string) string {
  slash := strings.LastIndex(s, "/") // -1 if "/" is not found
  s = s[slash + 1:]
  if dot := strings.LastIndex(s, "."); dot >= 0 {
    s = s[:dot]
  }
  return s
}
