package main
import "fmt"


func main()  {

    var a uint = 60
    var b uint = 13
    var c int = 4
    var ptr *int
    fmt.Printf("null ptr  => %d\n", ptr)
    // fmt.Printf("null ptr v=> %d\n", *ptr)
    fmt.Printf("bit and   => %d\n", a&b)
    fmt.Printf("bit or    => %d\n", a|b)
    fmt.Printf("bit not   => %d\n", a^b)
    fmt.Printf("bit left  => %d\n", a<<2)
    fmt.Printf("bit right => %d\n", a>>2)
    ptr = &c
    fmt.Printf("pt     r  => %d\n", ptr)
    fmt.Printf("pt v   r  => %d\n", *ptr)
    fmt.Printf("c address => %d\n", &c)
    fmt.Printf("c address => %d\n", c)
    c = 12
    fmt.Printf("pt     r  => %d\n", ptr)
    fmt.Printf("pt v   r  => %d\n", *ptr)
    fmt.Printf("c address => %d\n", &c)
    fmt.Printf("c address => %d\n", c)
}
