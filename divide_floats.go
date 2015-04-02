package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  a, _ := strconv.ParseFloat(os.Args[1], 64)
  b, _ := strconv.ParseFloat(os.Args[2], 64)
  prec := 2
  if len(os.Args) > 3 {
   prec, _ = strconv.Atoi(os.Args[3])
  }
  result := strconv.FormatFloat((a / b), 'f', prec, 64)
  fmt.Println("{\"result\":\""+result+"\"}")
}
