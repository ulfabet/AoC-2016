package main

import "fmt"

//const input = 5
const input = 3017957

func main() {
    var left int
    elves := make([]int, input)
    presents := make([]int, input)
    for i := range presents {
        presents[i] = 1
    }

    for i := 0; ; i = (i+1)%len(elves) {
        if presents[i] == 0 { continue }
        for j := 1; j < len(presents); j++ {
            left = presents[(i+j)%len(presents)]
            if left == 0 { continue }
            presents[i] += left
            presents[(i+j)%len(presents)] = 0
            //fmt.Println("Elf", i, "steals presents from", (i+j)%len(presents))
            break
        }
        if left == 0 {
            fmt.Println("Elf that gets all the presents:", i+1)
            break
        }
    }
}

