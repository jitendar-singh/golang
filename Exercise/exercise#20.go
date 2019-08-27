package main
import `fmt`

func main()  {

  s := struct{
    name string
    friends map[string]int
    skills []string
  }{
    name: `Jitendar Singh`,
    friends: map[string]int{
      `Monalisa Sarangi`: 7019289339,
      `Sandip Singh`: 9937003387,
      `Surjit Kaur`: 9853498534,
    },
    skills: []string{
      `golang`,
      `python`,
      `apex`,
      `Information Security`,
      `sales`,
      `Marketing`,
      `cooking`,
      `laundry`,
    },
  }

    fmt.Println(s.name)
    fmt.Println(`Friends List`)

    for k, fv := range s.friends{
      fmt.Println(k, fv)
    }

    fmt.Println(`Skills`)
    for _, sv := range s.skills{
      fmt.Println(sv)
    }
}
