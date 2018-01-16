package main

import "fmt"

//const input = 5
//const input = 6
const input = 3017957

var skip [input]bool

func advance(x, dx int) int {
    for i := 0; i < dx; i++ {
        x = (x+1)%input
        for skip[x] == true { x = (x+1)%input }
    }
    return x
}

func main() {

    //var i,j,left int
    j := input/2
    left := input

    //for i != j {
    for left > 1 {
        //fmt.Println("Elf", i+1, "steals presents from", j+1)
        skip[j] = true
        left--
        //i = advance(i, 1)
        j = advance(j, 2-left&1)
    }
    fmt.Println("Elf that gets all the presents:", j+1)
}

