package main
import `fmt`

type vehicle struct{
  doors int
  color string
}

type truck struct{
  vehicle
  fourWheel bool
}

type sedan struct{
  vehicle
  luxury bool
}

func main(){

  t1 := truck{
    vehicle: vehicle{
      doors: 2,
      color: `Red`,
    },
    fourWheel: false,
  }

 s1 := sedan{
   vehicle: vehicle{
     doors: 4,
     color: `Black`,
   },
   luxury: true,
 }

 fmt.Println(`Details of truck`,t1)
 fmt.Println(`Details of sedan`,s1)

 fmt.Println(`Truck doors:`,t1.doors)
 fmt.Println(`Truck color:`,t1.color)
 fmt.Println(`Truck has fourWheels:`,t1.fourWheel)


 fmt.Println(`Sedan doors:`,s1.doors)
 fmt.Println(`Sedan color:`,s1.color)
 fmt.Println(`Sedan is luxury:`,s1.luxury)

}
