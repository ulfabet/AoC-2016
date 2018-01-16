package main

import (
    "fmt"
    //"log"
    "os"
    "bufio"
)

// scan default or open paren
// scan first number or x
// scan second number or close paren
// scan characters to duplicate

var state int
var memory string
var output int
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
        output += len(memory)
        //log.Println(output, memory)
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
        tmp := memory[:]
        r := repeat
        for ;r > 0; r-- {
            process_line(tmp)
        }
        memory = ""
        state = 0
    }
}

func process_line(line string) {
    memory = ""
    state = 0
    for _, v := range line {
        process_character(string(v))
    }
    output += len(memory)
    //log.Println(output, memory)
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        for _, v := range s.Text() {
            process_character(string(v))
        }
    }
    output += len(memory)
    //log.Println(output, memory)
    fmt.Println("output", output)
}

