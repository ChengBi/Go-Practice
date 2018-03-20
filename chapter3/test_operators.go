package main
import "fmt"

func main() {
    var a int = 10
    var b int = 21

    fmt.Println(a - b)
    fmt.Println(a + b)
    fmt.Println(a * b)
    fmt.Println(a / b)
    fmt.Println(a % b)

    fmt.Println("----------------------------")
    fmt.Println(a)
    fmt.Println(&a)
    fmt.Println(*&a)
}
