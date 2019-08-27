package main
import "fmt"

type person struct{
  name string
  age int
  splz string
}
type assisn struct{
  person
  ltk bool
}

func main(){

    a1 := assisn{
      person: person{
        name: `James Bond`,
        age: 27,
        splz: `Feild Agent`,
      },
      ltk: true,
    }

    p1 := person{
      name: `Jitendar Singh`,
      age: 30,
      splz: `IT`,
    }

    fmt.Println(`Person Details with "Person Struct "`,p1)

    // fmt.Println(a1)
    fmt.Println(`Name: `,a1.name)
    fmt.Println(`Age: `,a1.age)
    fmt.Println(`Specilaization: `,a1.splz)
    fmt.Println(`License to Kill: `,a1.ltk)

}
