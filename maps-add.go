package main
import "fmt"

func main(){
  m := map[string]string{
    "Jitendar":"Singh",
    "Monalisa":"Sarangi",
    "Sunny":"Leone",
    "Larry":"Ellision",
  }

  if v, ok := m["Larry"];ok{
    fmt.Println("This is the if print ",v)
  }else{
    fmt.Println("Value not found")
  }
  m["Dinesh"] = "Swain"
  fmt.Println("Elements of the map")
  for k, v := range m {
    fmt.Println(k, v)
  }
}
