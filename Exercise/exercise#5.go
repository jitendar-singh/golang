package main
import "fmt"

type mytype int
var x mytype
var y int

func main(){



  fmt.Printf("Type of x [%T] \n",x)
  fmt.Println("Value of x ",x)
  x = 42
  fmt.Println("New Value of x ",x)
  y = int(x)
  fmt.Println("Value of y ",y)
  fmt.Printf("Type of y [%T] \n",y)

}
