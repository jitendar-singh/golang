package main
import "fmt"

var (
  a,b int
)

func main(){

  a := 42
  b := 48

  for a <= b{
    if a == 48{
      fmt.Printf("%d equal to %d\n",b , a)
      a++
    }else{
      fmt.Printf("%d greater than %d\n",b , a)
      a++
    }
 }
  fmt.Println(a,"greater than",b)
}
