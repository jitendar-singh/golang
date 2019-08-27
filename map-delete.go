package main
import "fmt"

func main(){

  m := map[string]string{
    "Jitendar":"Singh",
    "Surjit":"Kaur",
    "Monalisa":"Sarangi",
  }

  fmt.Println("Elements of Map-")
  for k, v := range m {
    fmt.Println(k, v)
  }
  delete(m, "Jitendar")

  fmt.Println("Elements of Map after delete operation-\n")
  for k, v := range m {
    fmt.Println(k, v)
  }


}
