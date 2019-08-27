package main

import "fmt"

var y float32
func main(){
  x :=42

  fmt.Println(x,y)
  foo()
}
func foo(){
  fmt.Println(y)
}
