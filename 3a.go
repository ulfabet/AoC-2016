package main

import (
    "fmt"
    "os"
    "bufio"
//    "log"
//    "strings"
)

func process_line(line string) {
    //log.Println(line)
    var a, b, c int
    fmt.Sscan(line, &a, &b, &c)
    switch {
        case a >= b + c: { fmt.Println("Not a valid triangle:", a, b, c) }
        case b >= c + a: { fmt.Println("Not a valid triangle:", a, b, c) }
        case c >= a + b: { fmt.Println("Not a valid triangle:", a, b, c) }
        default: { fmt.Println("Valid triangle:", a, b, c) }
    }
}

func process_triangle(a, b, c int) {
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        process_line(s.Text())
    }
}

