package main

import "fmt"

//const input = 3017957
const input = 5

type Elf struct {
    number int
    p *Elf
    n *Elf
}

func main() {
    var elf, first, prev *Elf

    first = new(Elf)
    prev = first
    for i := 1; i < input; i++ {
        elf = new(Elf)
        //if i == input/2 { mid = elf }
        elf.number = i
        elf.p = prev
        elf.p.n = elf
        prev = elf
    }
    first.p = elf
    first.p.n = first

    /*
    for {
        mid.p.n = mid.n
        mid.n.p = mid.p
        mid = mid.n
    }
    */
    for i := 0; i < input; i++ {
        fmt.Println(first.number)
        first = first.p
    }

    fmt.Println("Elf that gets all the presents:", 42)
}
