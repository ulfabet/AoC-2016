package main

import "fmt"

const input = 5
//const input = 3017957

var skip [input]bool

func advance(x, dx int) int {
    for i := 0; i < dx; i++ {
        x = (x+1)%input
        for skip[x] == true {
            fmt.Println("skip", x+1)
            x = (x+1)%input
        }
    }
    return x
}

func main() {

    i := input/2
    left := input

    for left > 1 {
        fmt.Println("Steal from elf:", i+1)
        skip[i] = true
        left--
        i = advance(i, 2-left&1)
    }
    fmt.Println("Elf that gets all the presents:", i+1)
}
