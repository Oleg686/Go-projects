package main
import ("fmt";"os")

func main() {
    var chislo int
    fmt.Println("Введите число")
    fmt.Fscan(os.Stdin,&chislo)
    if (chislo % 2 == 0){
     fmt.Println(chislo,"= Четное") 
    } else{
        fmt.Println(chislo,"= Нечетное") 
    }
    
}
