package main
import `fmt`

type Person struct{
  firstName string
  lastName string
  flavour []string
}
func main(){

  p1 := Person{
      firstName: `Sunny`,
      lastName: `Singh`,
      flavour: []string{`Choco Almond`,`Butter Scotch`,`Vanilla`},
    }

  p2 := Person{
    firstName: `Monalisa`,
    lastName: `Sarangi`,
    flavour: []string{`Tender Coconut`,`Choco Almond`,`Kulfi(Rabadi)`},
  }

  fmt.Println(p1.firstName,`,`,p1.lastName,` likes one of these`)
  for i, v := range p1.flavour{
    fmt.Println(i+1, v)
  }

  fmt.Println(p2.firstName,`,`,p2.lastName,` likes one of these`)
  for i, v := range p2.flavour{
    fmt.Println(i+1, v)
  }
}
