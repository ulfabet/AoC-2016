package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

var prev Set
var current Set
var next Set
var level int
var label []string

type State [4]int
type Set map[State]bool

func (s Set) add(t State) {
    s[t] = true
}

func (s Set) contains(t State) bool {
    _, ok := s[t]
    return ok
}

func finished(t State) bool {
    if t[0] == 0 && t[1] == 0 && t[2] == 0 {
        fmt.Println("Finished in", level, "steps")
        return true
    }
    return false
}

func contains_generators(things int) bool {
    return things & 0xAAAAAA != 0
}

func compatible_generator(things, index int) bool {
    return things & (1<<uint(index-1)) != 0
}

// todo: memoize shift results?
func illegal(t State) bool {
    for _,v := range t {
        for j := 2; j < int(labels); j += 2 {
            if v&(1<<uint(j)) != 0 {
                if contains_generators(v) && !compatible_generator(v, j) {
                    return true
                }
            }
        }
    }
    return false
}

func old_illegal(t State) bool {
    n := 0
    for _,v := range t {
        for j := 1; j < int(labels); j++ {
            if v&(1<<uint(j)) != 0 {
                if j&1 == 1 { // generator
                    n += 1
                } else { // microchip
                    n -= 1
                    if contains_generators(v) && !compatible_generator(v, j) {
                        return true
                    }
                }
            }
        }
        if n > 2 {
            //fmt.Println("illegal: three or more microchips above corresponding generators")
            return true
        }
    }
    return false
}

const labels uint = 15
//const labels uint = 5

func initialize() {
    current = make(Set)
    prev = current
    label = strings.Fields("ee sg sm pg pm tg tm rg rm cg cm eg em dg dm")
    //label = strings.Fields("ee hg hm lg lm")
}

// todo: remove need for var u State
func next_states(t State) Set {
    var elevator int
    for elevator = range t {
        if t[elevator]&1 == 1 { break }
    }
    s := make(Set)
    for i := 2; i <= 1<<labels; i = i<<1 {
        a := t[elevator]&i
        if a == 0 { continue }
        for j := i; j <= 1<<labels; j = j<<1 {
            b := t[elevator]&j
            if b == 0 { continue }
            c := a|b|1
            if elevator < 3 {
                var u State
                u[0], u[1], u[2], u[3] = t[0], t[1], t[2], t[3]
                u[elevator] -= c
                u[elevator+1] += c
                s.add(u)
                //fmt.Printf(" up %x\n", c)
            }
            if elevator > 0 {
                var u State
                u[0], u[1], u[2], u[3] = t[0], t[1], t[2], t[3]
                u[elevator] -= c
                u[elevator-1] += c
                s.add(u)
                //fmt.Printf(" down %x\n", c)
            }
        }
    }
    return s
}

func show(t State) {
    for i, v := range t {
        fmt.Print(i, " ")
        for j, w := range label {
            if v & (1<<uint(j)) != 0 {
                fmt.Print(w, " ")
            } else {
                fmt.Print(".. ")
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

func advance() bool {
    level++
    fmt.Println("level", level, len(current))
    //if level > 50 { return false }
    next = make(Set)
    for a := range current {
        for b := range next_states(a) {
            if prev.contains(b) { continue }
            if current.contains(b) { continue }
            //if next.contains(b) { continue } // not necessary
            if illegal(b) { continue }
            //show(b)
            if finished(b) { return false }
            next.add(b)
        }
    }
    prev = current
    current = next
    return true
}

type State struct

func main() {
    var t State
    initialize()
    s := bufio.NewScanner(os.Stdin)
    n := 0
    for s.Scan() {
        line := s.Text()
        fields := strings.Fields(line[:len(line)-1])
        for i := 0; i < len(fields); i++ {
            switch fields[i] {
            case "generator":
                fallthrough
            case "generator,":
                fallthrough
            case "microchip":
                fallthrough
            case "microchip,":
                thing := fmt.Sprintf("%c%c", fields[i-1][0], fields[i][0])
                for j, w := range label {
                    if w == thing {
                        t[n] = t[n]|(1<<uint(j))
                    }
                }
            }
        }
        n++
    }
    t[0] = t[0]|1 // elevator
    //show(t)
    current.add(t)
    for advance() {}
}

