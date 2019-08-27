package main
import "fmt"

var (
  a,b int
  res bool
)
func main(){
  a := 10
  b := -10

  fmt.Println(a==b)
  fmt.Println(a<=b)
  fmt.Println(a>=b)
  fmt.Println(a!=b)
  fmt.Println(a<b)
  fmt.Println(a>b)

}
