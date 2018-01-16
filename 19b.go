package main

import "fmt"

//const input = 5
const input = 3017957

func main() {
    elves := make([]int, input)
    numbers := make([]int, input)

    for i := range elves {
        elves[i] = 1
        numbers[i] = i+1
    }

    var i,j int
    for i = 0; len(elves) > 1; i = (i+1)%len(elves) {
        j = (i+len(elves)/2)%len(elves)
        //fmt.Println("Elf", numbers[i], "steals presents from", numbers[j])
        elves[i] += elves[j]
        elves = append(elves[:j], elves[j+1:]...)
        numbers = append(numbers[:j], numbers[j+1:]...)
        if j < i { i-- }
    }
    fmt.Println("Elf that gets all the presents:", numbers[i])
}

