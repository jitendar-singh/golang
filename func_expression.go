package main
import `fmt`

func main(){

  func1 := func(){
    fmt.Println(`This is my first function expression`)
  }

  fucn2 := func(name string){
    fmt.Println(`Welcome `,name, `to the world of function expression`)
  }

  func1()
  fucn2(`Jitendar Singh`)

  fmt.Printf("\nfunc2 is of type: %T\n ",fucn2)
}
