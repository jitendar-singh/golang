package main
import `fmt`

type Person struct{
  first string
  last string
}

type SecretAgent struct{
  Person
  ltk bool
}

func main(){

  sa1 := SecretAgent{
    Person: Person{
      first: `Jitendar`,
      last: `Singh`,
    },
    ltk: false,
  }

  sa2 := SecretAgent{
    Person: Person{
      first: `Monalisa`,
      last: `Sarangi`,
    },
    ltk: true,
  }
  sa1.agent()
  sa2.agent()
  
//  calling agent() directly will raise an undefined error
}

func (r SecretAgent) agent(){
  fmt.Println(`Welcome to MI6 Database`)
  fmt.Println(r.last, r.first)
  fmt.Println(`Your ltk is `, r.ltk)
}
