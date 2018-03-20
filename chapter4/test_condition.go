package main
import "fmt"


func main() {

    var a int = 1
    var b int = 2
    var c string = "1234"
    var d string = "132"
    // var e bool = true
    // var f bool = false

    if (a > b) {
        fmt.Printf("%d is bigger!\n", a)
    } else {
        fmt.Printf("%d is bigger!\n", b)
    }

    if (c > d) {
        fmt.Printf("%s is bigger!\n", c)
    } else {
        fmt.Printf("%s is bigger!\n", d)
    }

    // if (e > f) {
    //     fmt.Printf("%b is bigger!", e)
    // } else {
    //     fmt.Printf("%b is bigger!", f)
    // }

    for  i := 0; i < 10; i ++{
        switch i {
        case 0:
            fmt.Println(i)
        case 1:
            fmt.Println(i)
        case 2:
            fmt.Println(i)
        case 3:
            fmt.Println(i)
        case 4:
            fmt.Println(i)
        }
    }

    for i := 0; i < 10; i ++ {
        var c1, c2, c3 chan int
        var i1 = 1
        var i2 = 2
        select {
           case i1 = <-c1:
              fmt.Printf("received ", i1, " from c1\n")
           case c2 <- i2:
              fmt.Printf("sent ", i2, " to c2\n")
           case i3, ok := (<-c3):  // same as: i3, ok := <-c3
              if ok {
                 fmt.Printf("received ", i3, " from c3\n")
              } else {
                 fmt.Printf("c3 is closed\n")
              }
           default:
              fmt.Printf("no communication\n")
        }
    }

    var j int = 1
    var k *int = &j
    // var l = &k
    var l **int = &k

    fmt.Println(j)
    fmt.Println(k)
    fmt.Printf("%T\n", l)
    fmt.Println(*l)
    fmt.Println(**l)

}
