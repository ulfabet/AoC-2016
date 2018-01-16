package main

import (
    "fmt"
    "os"
    "bufio"
//    "log"
//    "strings"
)

func process_line(line string) (a,b,c int) {
    fmt.Sscan(line, &a, &b, &c)
    return
}

func process_triangle(a, b, c int) {
    switch {
        case a >= b + c: { fmt.Println("Not a valid triangle:", a, b, c) }
        case b >= c + a: { fmt.Println("Not a valid triangle:", a, b, c) }
        case c >= a + b: { fmt.Println("Not a valid triangle:", a, b, c) }
        default: { fmt.Println("Valid triangle:", a, b, c) }
    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        a, d, g := process_line(s.Text())
        s.Scan()
        b, e, h := process_line(s.Text())
        s.Scan()
        c, f, i := process_line(s.Text())
        process_triangle(a, b, c)
        process_triangle(d, e, f)
        process_triangle(g, h, i)
    }
}

