package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

var elevator int
var floor [4]int
var history map[int64]uint8
var label []string
var steps uint8
var solutions int

const labels uint = 15
//const labels uint = 5

func myinit() {
    history = make(map[int64]uint8)
    label = strings.Fields("ee sg sm pg pm tg tm rg rm cg cm eg em dg dm")
    //label = strings.Fields("ee hg hm lg lm")
    if len(label) != int(labels) { fmt.Println("Error: len(label)") } 
    floor[elevator] = 1 // elevator
}

func is_finished() bool {
    if floor[0] == 0 && floor[1] == 0 && floor[2] == 0 {
        fmt.Println("Finished in", steps, "steps")
        solutions++
        return true
    } else {
        return false
    }
}

func in_history() bool {
    var state int64
    state = int64(floor[3])<<48 | int64(floor[2])<<32 | int64(floor[1])<<16 | int64(floor[0])
    n, ok := history[state]
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

func add_history(n uint8) {
    var state int64
    state = int64(floor[3])<<48 | int64(floor[2])<<32 | int64(floor[1])<<16 | int64(floor[0])
    history[state] = n
}

func is_microchip(s string) bool {
    return s[1] == 'm'
}

func contains_generators(things int) bool {
    return things & 0xAAAAAA != 0
}

func compatible_generator(things, index int) bool {
    return things & (1<<uint(index-1)) != 0
}

func odd(n int) bool {
    return n&1 == 1
}
func even(n int) bool {
    return n&1 == 0
}

/*
func new_is_illegal() bool {
    for _,v := range floor {
        for j,w := range bits {
            if v&w != 0 {
                if is_microchip( ...
            }
        }
    }
}
*/

func is_illegal() bool {
    /*
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
    */
    n := 0
    for _,v := range floor {
        for j := 1; j < int(labels); j++ {
            if v&(1<<uint(j)) != 0 {
                if odd(j) { // generator
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
    if in_history() {
        return true
    }
    add_history(steps)
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

/*
1 => 1
2 => 2+1
3 => 3+3
4 => 4+6   (1 2 3 4)
5 => 5+10  (1 2 3 4 5)
6 => 6+(6 over 2)=6+15
*/

func combinations() []int {
    var r []int
    things := floor[elevator]

    i := 2
    for i <= 1<<labels {
        j := i
        for j <= 1<<labels {
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

/*
func ordered_combinations() []int {
    var r []int
    things := floor[elevator]

    var i,j int
    // generators only
    i = 2
    for i <= 1<<labels {
        j = i
        for j <= 1<<labels {
            a := things&(i|j)
            if a != 0 {
                r = append(r, a|1)
            }
            j = j<<2
        }
        i = i<<2
    }
    // microchips only
    i = 4
    for i <= 1<<labels {
        j = i
        for j <= 1<<labels {
            a := things&(i|j)
            if a != 0 {
                r = append(r, a|1)
            }
            j = j<<2
        }
        i = i<<2
    }
    // generators and microchips
    i = 2
    for i <= 1<<labels {
        j = i<<1
        for j <= 1<<labels {
            a := things&i
            b := things&j
            if a != 0 && b != 0 {
                r = append(r, a|b|1)
            }
            j = j<<2
        }
        i = i<<2
    }
    return r
}
*/

func move(direction, things int) {
    floor[elevator] -= things
    elevator += direction
    floor[elevator] += things
}

func dead_end() {
    // this only works if we did not hit the step limit...
    add_history(0)
}

func run() {
    if steps > 50 {
        return
    }
    if is_illegal() {
        return
    }
    fmt.Println(steps)
    show()
    if is_finished() {
        return
    }
    if elevator < 3 {
        i := 2
        for i <= 1<<labels {
            j := i
            for j <= 1<<labels {
                a := floor[elevator]&(i|j|1)
                if a != 1 {
                    move(1, a)
                    steps++
                    run()
                    move(-1, a)
                    steps--
                }
                j = j<<1
            }
            i = i<<1
        }
    }
    if elevator > 0 {
        i := 2
        for i <= 1<<labels {
            j := i
            for j <= 1<<labels {
                a := floor[elevator]&(i|j|1)
                if a != 1 {
                    move(-1, a)
                    steps++
                    run()
                    move(1, a)
                    steps--
                }
                j = j<<1
            }
            i = i<<1
        }
    }
    //dead_end()
    /*
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
    */
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

