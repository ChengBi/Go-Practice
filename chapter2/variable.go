package main
import "fmt"

var x, y int
var (
    a int
    b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"

func main(){
    g, h := "g", "h"
    fmt.Println(x, y, a, b, c, d, e, f, g, h)
    i, j := 123, 456
    fmt.Println(i, j)
    i = j // only copy values instead of reference
    fmt.Println(i, j)
    fmt.Println(&i, &j)

    j += 1
    fmt.Println(i, j)
    fmt.Println("--------------------------------")
    fmt.Println("address in stack")
    fmt.Println("x: <int>   ", x, &x)
    fmt.Println("y: <int>   ", y, &y)
    fmt.Println("a: <int>   ", a, &a)
    fmt.Println("b: <bool>  ", b, &b)
    fmt.Println("c: <int>   ", c, &c)
    fmt.Println("d: <int>   ", d, &d)
    fmt.Println("e: <int>   ", e, &e)
    fmt.Println("f: <string>", f, &f)
    fmt.Println("g: <string>", g, &g)
    fmt.Println("h: <string>", h, &h)
    fmt.Println("i: <int>   ", i, &i)
    fmt.Println("j: <int>   ", j, &j)

}
