package main
import "fmt"

func main(){
  m := map[string][]string{
    "bond_james": []string{"Shaken, not stirred","Matinis","women"},
    "moneypenny_miss": []string{`James Bond`,`Literature`,`Computer Science`},
    "no_dr": []string{`Being Evil`, `Ice Cream`, `Sunsets`},
  }
  fmt.Println(m)
}
