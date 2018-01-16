package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func atoi(s string) (i int) {
    fmt.Sscan(s, &i)
    return
}

//---- 
type Device interface {
}

//---- 
type Generator struct {
    material string
}

//---- 
type Microchip struct {
    material string
}

//---- 
type Floor struct {
    number int
    contents []*Device
}

//----
var floors [4]Floor
var solution []string

/*
Build a tree of all possible moves?

solution = ["up/down devices" ...]
answer = len(solution)

select some thing(s) from the first floor
    that can be moved to the second floor
    if not possible => give up
select some thing(s) from the second floor
    that can be moved to the third floor
    if not possible => move down (how far?)
select some thing(s) from the third floor
    that can be moved to the fourth floor
    if not possible => move down (how far?)
*/

func foo() {
    for i := 0; i < 4; i++ {
        f = new(Floor)
        f.number = i+1
        floors[i] = f
    }
    n
    i:=
    for {
        // if destination floor safe for devices

        var current, above, below *Floor
        current := floor[i]
        if i < 3 { above = floor[i+1] }
        if i > 0 { below = floor[i-1] }
        elevator.move(floor, devices)

    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    floor := 1
    for s.Scan() {
        line := s.Text()
        fields := strings.Fields(line[:len(line)-1])
        for len(fields) > 0 {
            switch fields[len(fields)-1] {
                case "generator": {
                    fmt.Println(floor, fields[len(fields)-2:])
                }
                case "microchip": {
                    fmt.Println(floor, fields[len(fields)-2:])
                }
            }
            fields = fields[:len(fields)-1]
        }
        floor++
    }
}

