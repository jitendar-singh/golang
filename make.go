package main
import "fmt"

func main(){
  x := make([]int, 10, 12)

  fmt.Println(len(x))
  fmt.Println(cap(x))
  x = append(x,12)
  fmt.Println(len(x))
  fmt.Println(cap(x))
  x = append(x,13)
  fmt.Println(len(x))
  fmt.Println(cap(x))
  x = append(x,14)
  fmt.Println(len(x))
  fmt.Println(cap(x))
  x = append(x,14)
  fmt.Println(len(x)) 
  fmt.Println(cap(x))
  fmt.Println(x)

}
