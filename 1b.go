package main

import (
    "fmt"
)

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
    var s string
    var n int
    var m [1000000]int
    var i int
    var x, y int

    /*
    for j := 0; j < 500000; j++ {
        m[j] = 0
    }
    fmt.Println("debug")
    */

    /*
    for j := 0; j < 1000; j++ {
        for k := 0; k < 1000; k++ {
            m[j][k] = 0
        }
    }
    */

    for {
        _, err := fmt.Scanf("%1s%d", &s, &n)
        if err != nil {
            fmt.Println(err)
            break
        }
        switch s {
            case "L": {
                i = (i-1) & 3
                fmt.Println("Left", n, i)
            }
            case "R": {
                i = (i+1) & 3
                fmt.Println("Right", n, i)
            }
            default: {
                fmt.Println("Error")
                break
            }
        }
        for n > 0 {
            n--
            switch i {
                case 0: { y++ }
                case 1: { x++ }
                case 2: { y-- }
                case 3: { x-- }
            }
            num := m[(500+x)+1000*(500+y)]
            //fmt.Println("debug", x, y)
            if num > 0 {
                fmt.Println("Here!", x, y, abs(x)+abs(y))
                return
            }
            m[(500+x)+1000*(500+y)] += 1
        }
    }
}

