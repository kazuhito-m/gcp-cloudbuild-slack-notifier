package main

import (
  "fmt"
  "github.com/satori/go.uuid"
)
 
func main() {
    fmt.Printf("%v\n", uuid.Must(uuid.NewV4()));
}