package main
import "fmt"

func main(){

  var x= []int{4, 5, 7, 8, 42}
  x = append(x,7,8,9,10)
  fmt.Println(x)

  x = append(x[:3],x[5:]...)
  fmt.Println(x)

}
