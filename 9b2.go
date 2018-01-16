package main

import (
    "fmt"
    "log"
    "os"
    "bufio"
)

// scan default or open paren
// scan first number or x
// scan second number or close paren
// scan characters to duplicate

var state int
var memory string
var output uint64
var counter int
var repeat int

func process_character(c string) {
    //log.Println("c", c)
    switch state {
        case 0: { scan_default_or_open_paren(c) }
        case 1: { scan_first_number_or_x(c) }
        case 2: { scan_second_number_or_close_paren(c) }
        case 3: { scan_characters_to_repeat(c) }
    }
    update_cstack()
}

func rstack_product() int {
    p:= 1
    for _, v := range rstack { p *= v }
    return p
}

func update_output() {
    if len(rstack) > 0 {
        output += uint64(rstack_product() * len(memory))
    } else {
        output += uint64(len(memory))
    }
    //log.Println("rstack", rstack)
    //log.Println("output", output, memory)
}

func update_cstack() {
    for i := len(cstack)-1; i >= 0; i-- {
        cstack[i]--
        if cstack[i] == 0 {
            if len(memory) > 0 {
                update_output()
                memory = ""
                state = 0
            }
            cstack = cstack[:i]
            //log.Println("pop cstack", cstack)
            rstack = rstack[:i]
            //log.Println("pop rstack", rstack)
        }
    }
}

func scan_default_or_open_paren(c string) {
    if c == "(" {
        update_output()
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
        cstack = append(cstack, counter+1)
        //log.Println("push cstack", cstack)
        rstack = append(rstack, repeat)
        //log.Println("push rstack", rstack)
        memory = ""
        state = 0
    } else {
        memory += c
    }
}

var rstack []int
var cstack []int

// not in use
func scan_characters_to_repeat(c string) {
    log.Println("error")
    os.Exit(1)
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

