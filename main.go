package main

import (
  "astro/compiler/parser"
  "io/ioutil"
  "fmt"
  "os"
  "log"
)

func main(){
  if len(os.Args) < 2 {
    fmt.Printf("%v FILE\n", os.Args[0])
    os.Exit(1)
  }

  buffer, err := ioutil.ReadFile("doc/try.fxl")
  if err != nil {
    log.Fatal(err)
  }

  parser.Parse(string(buffer))
}
