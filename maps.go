package main
import "fmt"

func main(){

  m := map[string]int{
    "James":32,
    "Miss Moneypenny":27,
    "Sunny":30,
    "Monalisa":28,
  }
  fmt.Println(m)
  fmt.Println(m["James"])
  fmt.Println(m["Sunny"])

  if v, ok := m["Monalisa"]; ok {
    fmt.Println("This is the if Print",v)
  }

  v, ok := m["Sunnys"]
  fmt.Println(v)
  fmt.Println(ok)
}
