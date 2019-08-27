package main
import "fmt"

func main(){

  states := make([]string, 15, 30)
  states = []string{"Odisha","Bihar","Assam","Tripura","Meghalaya","J&K","Punjab"}
  fmt.Println(len(states))
  fmt.Println(cap(states))

  for i :=0; i < len(states); i++{
    fmt.Println(i, states[i])
  }


}
