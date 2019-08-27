package main

import "fmt"

func main(){
  x := []int{1, 2, 3, 4, 5, 6, 42}
  fmt.Println("Length of x is ",len(x))
  for i:=0; i<len(x); i++{
    fmt.Printf("x[%d]=%d \n",i,x[i])
  }
  fmt.Println("Using range now")

  for i, v := range x{
    fmt.Println("[",i,"]=",v)
  }
}
