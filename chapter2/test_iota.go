package main
import "fmt"

const (
    a = 1<<iota//0
    b = 3<<iota//1
    c = 5<<iota//2
    d//3
)




func main()  {
    fmt.Println(a, b, c, d)
}
