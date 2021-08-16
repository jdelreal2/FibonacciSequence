// GoLang - Fibonacci Sequence

package main
import "fmt"

func main(){
    fmt.Println("Fibonacci Sequence")
    var finalNumber int
    num1:=0
    num2:=1
    nextNum:=0
    fmt.Println("How many terms should be printed: ")
    fmt.Scanln(&finalNumber)
    fmt.Println("Requested Fibonacci series: ")
    for i:=1;i<=finalNumber;i++ {
        if(i==1){
            fmt.Println(" ", num1)
            continue
        }
        if(i==2){
            fmt.Println(" ", num2)
            continue
        }
        nextNum = num1 + num2
        num1 = num2
        num2 = nextNum
        fmt.Println(" ", nextNum)
    }
}
