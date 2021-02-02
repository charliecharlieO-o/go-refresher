package main

import "fmt"

func main() {
  // Built in append function demo
  var runes []rune
  for _, r := range "Hello, World" {
    runes = append(runes, r)
  }
  fmt.Printf("%q\n", runes)

  // A closer look at what the append OP does
  // each change in capacity indicates an allocation and a copy
  var x, y []int
  for i := 0; i < 10; i++ {
    y = appendInt(x, i)
    fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
    // updating the slice variable is required for any function that may change
    // length or capacity of a slice or make it refer to a different underlying array
    // in this sense slices are not "pure" reference types
    x = y
  }
}

// NOTE: the built in append function may use a more sophisticated growth strategy
func appendInt(x []int, y int) []int {
  var z []int
  zlen := len(x) + 1
  if zlen <= cap(x) {
    // There is room to grow. Extend the slice
    z = x[:zlen]
  } else {
    // There is insufficient space. Allocate a new array.
    // Grow by doubling, for amoritzed linear complexity
    zcap := zlen
    if zcap < 2*len(x) {
      zcap = 2 * len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x) // built in function
  }
  z[len(x)] = y
  return z
}
