package main
import ("fmt";"os")

func main() {
    var chislo int
    fmt.Println("Введите число, которое делится на 3")
    fmt.Fscan(os.Stdin,&chislo)
    if (chislo % 3 == 0){
     fmt.Println(chislo,"= делится без остатка") 
    } else{
        fmt.Println(chislo,"= не делится без остатка") 
    }
    
}
