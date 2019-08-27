package main
import "fmt"

type mytype int

func main(){

  var x mytype

  fmt.Printf("Type of x [%T] \n",x)
  fmt.Println("Value of x ",x)
  x = 42
  fmt.Println("New Value of x ",x)

}
