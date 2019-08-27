package main
import "fmt"

type name string
var firstname name
var lastname name

func main(){
  fmt.Println("Enter firstname:")
  fmt.Scanf("%s",&firstname)
  fmt.Println("Enter lastname:")
  fmt.Scanf("%s",&lastname)
  fmt.Printf("Welcome %s %s\n ",firstname,lastname)
  fmt.Printf("%T \t %T\n",firstname,lastname)
}
