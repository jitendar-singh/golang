package main
import "fmt"

func main(){
  var x [5]int
  fmt.Println(x)
  x[3] = 42

  for i := 0; i <= len(x) ; i++{
    fmt.Println(i)
  }
}
