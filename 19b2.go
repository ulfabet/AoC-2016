package main

import "fmt"

const input = 5
//const input = 3017957

func main() {
    elves := make([]int, input)

    next := make([]int, input)
    target := make([]int, input)

    skip := 0
    n := len(elves)

    for i := range elves {
        elves[i] = (i+n/2)%n
    }

    var i,j,k int
    for i = 0; n-skip > 1; {
        //fmt.Println("skip", skip)
        //if skip % 100 == 0 { fmt.Println("skip", skip) }
        j = elves[i]
        fmt.Println("Elf", i+1, "steals presents from", j+1)
        elves[i] = elves




        prev := j
        for k = (n-skip)/2; k > 0; k-- {
            prev = j
            j = (j+elves[j])%n
            //fmt.Println(i, j, k)
        }
        //fmt.Println("Elf", i+1, "steals presents from", j+1)
        elves[prev] += 1
        skip++
        if skip+1 == n { break }
        i = (i+elves[i])%n
    }
    fmt.Println("Elf that gets all the presents:", i+1)
}

