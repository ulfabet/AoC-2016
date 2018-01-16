package main

import (
    "fmt"
    "os"
    "bufio"
)

// scan default or open paren
// scan first number or x
// scan second number or close paren
// scan characters to repeat

var state int
var memory string
var output uint64
var repeat int
var rstack []int
var counter int
var cstack []int


func process_character(c string) {
    switch state {
        case 0: { scan_default_or_open_paren(c) }
        case 1: { scan_first_number_or_x(c) }
        case 2: { scan_second_number_or_close_paren(c) }
        //case 3: { scan_characters_to_repeat(c) }
    }
    update_cstack()
}

func rstack_product() int {
    p := 1
    for _, v := range rstack { p *= v }
    return p
}

func update_output() {
    output += uint64(rstack_product() * len(memory))
    memory = ""
}

func update_cstack() {
    for i := len(cstack)-1; i >= 0; i-- {
        cstack[i]--
        if cstack[i] == 0 {
            update_output()
            cstack = cstack[:i]
            rstack = rstack[:i]
        }
    }
}

func scan_default_or_open_paren(c string) {
    if c == "(" {
        update_output()
        state = 1
    } else {
        memory += c
    }
}

func scan_first_number_or_x(c string) {
    if c == "x" {
        fmt.Sscan(memory, &counter)
        memory = ""
        state = 2
    } else {
        memory += c
    }
}

func scan_second_number_or_close_paren(c string) {
    if c == ")" {
        fmt.Sscan(memory, &repeat)
        cstack = append(cstack, counter+1)
        rstack = append(rstack, repeat)
        memory = ""
        state = 0
    } else {
        memory += c
    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        for _, v := range s.Text() {
            process_character(string(v))
        }
    }
    update_output()
    fmt.Println("output", output)
}

