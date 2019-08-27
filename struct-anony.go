package main
import "fmt"

func main(){

  p1 := struct{
    name string
    age int
    state string
  }{
    name: `Jitendar Singh`,
    age: 29,
    // state: `Odisha`,
  }
  fmt.Println(p1)
}
