package main

import "fmt"

func main() {
  var a [3]int
  fmt.Println(a[0])
  fmt.Println(a[len(a) - 1])

  // Print the indeces and elements
  for i, v := range a {
    fmt.Printf("%d %d\n", i, v)
  }

  // Print the elements only
  for _, v := range a {
    fmt.Printf("%d\n", v)
  }

  // Array literals
  var q [3]int = [3]int{1, 2, 3}
  var r [3]int = [3]int{1, 2}
  fmt.Println(q)
  fmt.Println(r[2])

  // Ellipsis
  x := [...]int{1, 2, 3}
  fmt.Printf("%T\n", x)
}
