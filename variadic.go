package main
import `fmt`

func main(){

  result := sum("Jitendar",89, 33, 4, 5, 77)
  fmt.Println(`Result = `,result)
}
func sum(name string, xi ...int) int {

  fmt.Println(`Hello `, name)
  fmt.Printf("Length of slice %d\n", len(xi))
  result := 0
  for _, v := range xi {
    result += v
  }
  return result


}
