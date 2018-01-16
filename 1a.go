package main

import (
    "fmt"
)

func main() {
    var s string
    var n int
    var a [4]int
    var i int
    for {
        _, err := fmt.Scanf("%1s%d", &s, &n)
        if err != nil {
            fmt.Println(err)
            break
        }
        //fmt.Println(s, n)
        switch s {
            case "L": {
                i = (i-1) & 3
                fmt.Println("Left", n, i)
                a[i] = a[i] + n
            }
            case "R": {
                i = (i+1) & 3
                fmt.Println("Right", n, i)
                a[i] = a[i] + n
            }
            default: {
                fmt.Println("Error")
                break
            }
        }
    }
    y := a[0]-a[2]
    x := a[1]-a[3]
    fmt.Println(x, y, x+y)
}

