package main
import `fmt`

func main(){
  defer foo()
  defer bike()
   bar()
  defer car()
}

func foo(){
  fmt.Println("foo")
}

func bar(){
  fmt.Println("bar")
}

func car(){
  fmt.Println("car")
}

func bike(){
  fmt.Println("bike")
}
