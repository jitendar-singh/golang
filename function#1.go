package main
import `fmt`

func main(){

  xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  result := sum(xi ...)

  fmt.Println("\n Sum of first 10 +ve Integers is ", result)
}

func sum(xi ...int) int {
   fmt.Println("The first 10 +ve Integers are")
   var result int
   for _, v := range xi{
     fmt.Printf("%d \t",v)
     result += v
   }
   return result

}
