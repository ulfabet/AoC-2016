package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

var elevator int
var floor [4]int
var history map[[4]int]int
var label []string
var steps int

func myinit() {
    history = make(map[[4]int]int)
    label = strings.Fields("ee sg sm pg pm tg tm rg rm cg cm eg em dg dm hg hm lg lm")
    floor[elevator] = 1 // elevator
}

func is_finished() bool {
    if floor[0] == 0 && floor[1] == 0 && floor[2] == 0 {
        fmt.Println("Finished in", steps, "steps")
        return true
    } else {
        return false
    }
}

func in_history() bool {
    n, ok := history[floor]
    //fmt.Println(n, ok)
    if ok {
        if steps < n {
            //fmt.Println("found shorter path")
            return false
        } else {
            return true
        }
    } else {
        return false
    }
}

func add_history() {
    history[floor] = steps
}

func is_microchip(s string) bool {
    return s[1] == 'm'
}

func contains_generators(things int) bool {
    return things & 0xAAAAAA > 0
}

func compatible_generator(things, index int) bool {
    return things & (1<<uint(index-1)) > 0
}

func is_illegal() bool {
    if in_history() {
        //fmt.Println("in history")
        return true
    }
    add_history()
    for _,v := range floor {
        for j,w := range label {
            a := v&(1<<uint(j))
            if a == 0 { continue }
            if is_microchip(w) && contains_generators(v) && !compatible_generator(v, j) {
                //fmt.Println("illegal")
                return true
            }
        }
    }
    return false
}

func add(n int, thing string) {
    for i, v := range label {
        if v == thing {
            floor[n] = floor[n] | (1<<uint(i))
        }
    }
}

func show() {
    for i, v := range floor {
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

func combinations() []int {
    var r []int
    things := floor[elevator]

    i := 2
    for i <= 1<<18 {
        j := i
        for j <= 1<<18 {
            a := things&(i|j)
            if a != 0 {
                r = append(r, a|1)
            }
            j = j<<1
        }
        i = i<<1
    }
    return r
}

func move(direction, things int) {
    floor[elevator] -= things
    elevator += direction
    floor[elevator] += things
}

func run() {
    if is_illegal() {
        return
    }
    if steps > 50 {
        return
    }
    //fmt.Println(steps)
    //show()
    if is_finished() {
        return
    }
    for _,v := range combinations() {
        if elevator < 3 {
            //fmt.Println("move up from floor", elevator, ":", v)
            move(1, v)
            steps++
            run()
            move(-1, v)
            steps--
        }
        if elevator > 0 {
            //fmt.Println("move down from floor", elevator, ":", v)
            move(-1, v)
            steps++
            run()
            move(1, v)
            steps--
        }
    }
}

func main() {
    myinit()
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
                add(n, thing)
            }
        }
        n++
    }
    run()
}

