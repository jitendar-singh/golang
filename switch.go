package main
import "fmt"

func main(){

  var ch string

  fmt.Println("Enter a alphabet:")
  fmt.Scanf("%s",&ch)

  switch ch{
  case "a":
    fallthrough
  case "e":
    fallthrough
  case "i":
    fallthrough
  case "o":
    fallthrough
  case "u":
    fallthrough
  case "A":
    fallthrough
  case "E":
    fallthrough
  case "I":
    fallthrough
  case "O":
    fallthrough
  case "U":
    fmt.Println("VOWEL")
  default:
    fmt.Println("CONSONANT")
  }

}
