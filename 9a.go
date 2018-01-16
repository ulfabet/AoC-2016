package main

import (
    "fmt"
    "os"
    "bufio"
)

// scan default or open paren
// scan first number or x
// scan second number or close paren
// scan characters to duplicate

var state int
var memory string
var output string
var counter int
var repeat int

func process_character(c string) {
    switch state {
        case 0: { scan_default_or_open_paren(c) }
        case 1: { scan_first_number_or_x(c) }
        case 2: { scan_second_number_or_close_paren(c) }
        case 3: { scan_characters_to_repeat(c) }
    }
}

func scan_default_or_open_paren(c string) {
    if c == "(" {
        output += memory
        memory = ""
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
        memory = ""
        state = 3
    } else {
        memory += c
    }
}

func scan_characters_to_repeat(c string) {
    memory += c
    counter--
    if counter == 0 {
        for ;repeat > 0; repeat-- {
            output += memory
        }
        memory = ""
        state = 0
    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        for _, v := range s.Text() {
            process_character(string(v))
        }
    }
    output += memory
    fmt.Println("length", len(output))
    //fmt.Println("output", output)
}

