package main
import "fmt"

type tank struct {
    name string
    year int
}

func main() {
    fmt.Println("Бронетехника 2 мировой войны")
   var x = tank {"КВ",1939}
   fmt.Println(x.name, "= Название")
   fmt.Println(x.year, "= Год создания")
   
   fmt.Println("---------------------------------------------")
   
   var y = tank {"ИС",1944}
   fmt.Println(y.name, "= Название")
   fmt.Println(y.year, "= Год создания")
   fmt.Println("---------------------------------------------")
   if (x == y){
       fmt.Println("Два одинаковых танка")
   }else{
       fmt.Println("Разные модели")
   }
   
}
