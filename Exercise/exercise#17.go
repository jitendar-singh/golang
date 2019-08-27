package main
import "fmt"


type Person struct{
  first string
  last string
  age int
  ltk bool
  flavour []string
  }

func main(){

    p1 := Person{
    first: `Jitendar`,
    last: `Singh`,
    age: 27,
    ltk: true,
    flavour: []string{`Choco Almond`,`Butter Scotch`,`Vanilla`},
  }

  p2 := Person{
  first: `Monalisa`,
  last: `Sarangi`,
  age: 26,
  ltk: false,
  flavour: []string{`Tender Coconut`,`Choco Almond`,`Kulfi(Rabadi)`},
}
/* Use a map to store the struct and range through all the data in the struct */

m := map[string]Person{
  p1.last : p1,
  p2.last : p2,
}

  for _, v := range m{
    fmt.Println(v.first)
    fmt.Println(v.last)
    fmt.Println(v.age)
    fmt.Println(v.ltk)
    for i, v := range v.flavour{
      fmt.Println(i+1, v)
    }
    fmt.Println(`------------------`)
  }
}
