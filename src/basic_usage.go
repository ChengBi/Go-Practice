package main

import (
    "fmt"
    "math/rand"
    "math"
)

func add (x int, y int) int {
    return x + y
}

func add_1 (x, y int) int {
    return (add(x, y))
}

func swap (x, y int) (int, int) {
    return y, x
}

func neighbour(seed int) (x, y int) {
    x = seed - 1
    y = seed + 1
    return
}

var outer_list_b bool
var outer_list_i int
var outer_list_f float32
var outer_list_s string

var (
    max_int int = 1 << 32 -1
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // 这里开始就不能使用 v 了
    return lim
}

func main() {

    x := 20
    y := 34
    fmt.Println("Original data: ", x, y)
    fmt.Println("Print int value: ", rand.Intn(10))
    fmt.Println("Print pi: ", math.Pi)
    fmt.Println("add function result: ", x, y, " => ", add(x, y))
    fmt.Println("add_1 function result: ", x, y, " => ", add_1(x, y))
    a_swap, b_swap := swap(x, y)
    fmt.Println("swap function result: ", x, y, " => ", a_swap, b_swap)
    neighbour_x, neighbour_y := neighbour(x)
    fmt.Println("neighbour function: ", x, " => ", neighbour_x, neighbour_y)

    var inner_list_b bool
    var inner_list_i int
    var inner_list_f float32
    var inner_list_s string

    fmt.Println("outer list bool: ", outer_list_b)
    fmt.Println("outer list int: ", outer_list_i)
    fmt.Println("outer list float: ", outer_list_f)
    fmt.Println("outer list string: ", outer_list_s)
    fmt.Println("inner list bool: ", inner_list_b)
    fmt.Println("inner list int: ", inner_list_i)
    fmt.Println("inner list float: ", inner_list_f)
    fmt.Println("inner list string: ", inner_list_s)
    fmt.Println("max_int: ", max_int)

    var signed_int int = 0
    var usigned_int uint = 0
    signed_int = int(inner_list_f)
    fmt.Println("signed: ", signed_int, "usigned: ", usigned_int)
    fmt.Printf("sqrt function: %T , %v \n", math.Sqrt(4), math.Sqrt(4))
    fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
