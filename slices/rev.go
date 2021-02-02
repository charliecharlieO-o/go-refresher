package main

import "fmt"

func main() {
  // We are going to reverse an array
  a := [...]int{0, 1, 2, 3, 4, 5}
  reverse(a[:])
  fmt.Println(a)

  // Rotate by 2 positions
  reverse(a[:2])
  reverse(a[2:])
  fmt.Println(a)

  // NOTE: Slices don't have deep equality comparisons like arrays
  // NOTE: The only valid comparison is with nil
}

func reverse(s []int) {
  for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
    s[i], s[j] = s[j], s[i]
  }
}
