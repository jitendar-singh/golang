package main
import "fmt"

type person struct{
  first string
  last string
}
func main(){
  p1 := person{
    first: `James`,
    last: `Bond`,
  }

  p2 := person{
    first: `Jitendar`,
    last: `Singh`,
  }

  fmt.Println(p1)
  fmt.Println(p2)

  fmt.Println(p1.first)
  fmt.Println(p1.last)

  fmt.Println(p2.first)
}
