package main
import (
  `fmt`
  `time`
)


type Person struct{
  first string
  last string
  years_of_service int
  kills int
}

type SecretAgent struct{
  Person
  ltk bool
}

type human interface{
  database()
}

func main(){

  s1 := SecretAgent{
    Person: Person{
      first: `Jitendar`,
      last: `Singh`,
      years_of_service: 5,
      kills: 56,
    },
    ltk: true,
  }

  s2 := SecretAgent{
    Person: Person{
      first: `Sandip`,
      last: `Singh`,
      years_of_service: 12,
      kills: 190,
    },
    ltk: true,
  }

  s3 := SecretAgent{
    Person: Person{
      first: `Shaviraj`,
      last: `Singh`,
      years_of_service: 1,
      kills: 0,
    },
    ltk: false,
  }
 s1.database()
 s2.database()
 s3.database()
 fmt.Println(`------------------------------------`)
 fmt.Println(`Now lets fetch the method using Person`)

 p1 := Person{
   first: `Monalisa`,
   last: `Sarangi`,
   years_of_service: 5,
   kills: 100,
 }

p1.database()
database(p1)
database(s1)
database(s2)
database(s3)

fmt.Println(`Exiting the code please wait.........`)
time.Sleep(2 * time.Second)
}

func (r1 SecretAgent) database(){
  fmt.Println(`------------------------------------`)
  fmt.Println(`Welcome `, r1.last,`, `,r1.first)
  fmt.Println(`Please wait till we get your details`)
  time.Sleep(2 * time.Second)
  // fmt.Println(`------------------------------------`)
  fmt.Println(`Name: `,r1.first,` `,r1.last)
  fmt.Println(`Years of Service: `,r1.years_of_service)
  fmt.Println(`License to Kill: `,r1.ltk)
  fmt.Println(`Kills: `,r1.kills)

}

func (r2 Person) database(){
  fmt.Println(`------------------------------------`)
  fmt.Println(`Welcome `, r2.last,`, `,r2.first)
  fmt.Println(`Please wait till we get your Personal details`)
  time.Sleep(2 * time.Second)
  fmt.Println(`------------------------------------`)
  fmt.Println(`Name: `,r2.first,` `,r2.last)
  fmt.Println(`Years of Service: `,r2.years_of_service)
  // fmt.Println(`License to Kill: `,r1.ltk)
  fmt.Println(`Kills: `,r2.kills)
  fmt.Println(`------------------------------------`)

}

func database(h human){

  switch h.(type) {
  case Person:
    fmt.Println(`Hello I was passed into database in Interface `,h.(Person).first)
  case SecretAgent:
    fmt.Println(`Hello I was passed into database in Interface `,h.(SecretAgent).first)
  }
  fmt.Println(`Hello I was passed into database in Interface`,h)

}
