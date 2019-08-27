package main
import `fmt`

type Employee struct{
  firstName string
  lastName string
  empid int
  designation string
  ctc float32
  skills []string
}

func main(){

// Intialize the struct with Employee data

  emp1 := Employee{
    firstName: `Jitendar`,
    lastName: `Singh`,
    empid: 988698,
    designation: `Associate Customer Success Manager`,
    ctc: 977000.55,
    skills: []string{`python`,`golang`,`sql`,`aws`,`apex`,`docker`,`kubernetes`},
  }

  emp2 := Employee{
    firstName: `Monalisa`,
    lastName: `Sarangi`,
    empid: 003646,
    designation: `Information Security Advisor`,
    ctc: 825000.45,
    skills: []string{`ISO-27001`,`Risk Management`,`Vulnerablity Management`,`IBM AppScan`},
  }


  fmt.Println(`Employee Database using struct`)
  fmt.Println(`-----------------`)

  //Display records for 1st Employee


    fmt.Println(emp1.lastName,`,`,emp1.firstName)
    fmt.Printf(`%od`,emp1.empid)
    fmt.Println(``)
    fmt.Println(emp1.designation)
    fmt.Println(emp1.ctc)
    for _, v := range emp1.skills{
      fmt.Println(v)
    }
    fmt.Println(`---------------------`)

  //Display records for 2nd Employee


    fmt.Println(emp2.lastName,`,`,emp2.firstName)
    fmt.Printf(`%od`,emp2.empid)
    fmt.Println(``)
    fmt.Println(emp2.designation)
    fmt.Println(emp2.ctc)
    for _, v := range emp2.skills{
      fmt.Println(v)
    }
    fmt.Println(`---------------------`)

  //Dispaly records using map

    fmt.Println(`Employee Database using map`)
    fmt.Println(`-----------------`)

    m := map[string]Employee{
      emp1.firstName: emp1,
      emp2.firstName: emp2,
    }

    for _, v := range m{
      fmt.Println(v.firstName,`,`,v.lastName)
      fmt.Printf(`%od`,v.empid)
      fmt.Println(``)
      fmt.Println(v.designation)
      fmt.Println(v.ctc)

      for _, skv := range v.skills{
        fmt.Println(skv)
      }
      fmt.Println(`------------------------`)
    }


}
