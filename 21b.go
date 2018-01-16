package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

//const input = "abcde"
//const input = "decab"

//const input = "abcdefgh"
//const input = "ghfacdbe"
const input = "fbgdceah"

type Password struct {
    text []byte
}

func (p *Password) swap_position(a, b int) {
    p.text[a], p.text[b] = p.text[b], p.text[a]
}

func (p *Password) swap_letter(a, b byte) {
    for i := range p.text {
        switch p.text[i] {
            case a: p.text[i] = b
            case b: p.text[i] = a
        }
    }
}

func (p *Password) rotate(direction string, a int) {
    // flipped left and right for part 2
    n := len(p.text)
    a %= n
    switch direction {
    case "left":
        p.text = append(p.text[(n-a)%n:], p.text[:(n-a)%n]...)
    case "right":
        p.text = append(p.text[a%n:], p.text[:a%n]...)
    default:
        fmt.Println("rotate error", a)
    }
}

func (p *Password) index_of(a byte) int {
    for i := range p.text {
        if a == p.text[i] {
            return i
        }
    }
    return 999 // error
}

func (p *Password) rotate_based(a byte) {
    var i int
    i = len(p.text)-1
    if i >= 4 { p.rotate("right", 1) }
    p.rotate("right", 1)
    p.rotate("right", i)
    for {
        fmt.Println("debug", i, string(p.text))
        if p.index_of(a) == i {
            return
        }
        i--
        p.rotate("left", 1)
        if i == 3 { p.rotate("left", 1) }
    }
}

func (p *Password) reverse(a, b int) {
    if a > b { a, b = b, a }
    d := (b-a)/2
    for i := 0; i <= d; i++ {
        //fmt.Println(len(p.text), a+i, b-i)
        p.text[a+i], p.text[b-i] = p.text[b-i], p.text[a+i]
    }
}

func (p *Password) move(a, b int) {
    if a == b { return }
    a, b = b, a // part 2
    if a < b {
        d := b-a
        tmp := p.text[a]
        var i int
        for i = 0; i < d; i++ {
            p.text[a+i] = p.text[a+i+1]
        }
        if a+i != b { fmt.Println("error a < b:", a+i, b) }
        p.text[b] = tmp
    } else {
        d := a-b
        tmp := p.text[a]
        var i int
        for i = 0; i < d; i++ {
            p.text[a-i] = p.text[a-i-1]
        }
        if a-i != b { fmt.Println("error a >= b:", a-i, b) }
        p.text[b] = tmp
    }
}

type Stack []string

func (s Stack) push(a string) Stack {
    return append(s, a)
}

func (s Stack) pop() (Stack, string) {
    n := len(s)
    return  s[:n-1], s[n-1]
}

func main() {
    p := Password{[]byte(input)}
    s := bufio.NewScanner(os.Stdin)

    var lines Stack
    var line string

    for s.Scan() {
        lines = lines.push(s.Text())
    }

    for len(lines) > 0 {
        lines, line = lines.pop()
        f := strings.Fields(line)
        fmt.Println(f)
        switch f[0] {
        case "swap":
            switch f[1] {
            case "position":
                var a,b int
                fmt.Sscanf(f[2], "%d", &a)
                fmt.Sscanf(f[5], "%d", &b)
                p.swap_position(a, b)
            case "letter":
                p.swap_letter(f[2][0], f[5][0])
            }
        case "rotate":
            switch f[1] {
            case "based":
                p.rotate_based(f[6][0])
            default: // left or right
                var a int
                fmt.Sscanf(f[2], "%d", &a)
                p.rotate(f[1], a)
            }
        case "reverse":
            var a,b int
            fmt.Sscanf(f[2], "%d", &a)
            fmt.Sscanf(f[4], "%d", &b)
            p.reverse(a, b)
        case "move":
            var a,b int
            fmt.Sscanf(f[2], "%d", &a)
            fmt.Sscanf(f[5], "%d", &b)
            p.move(a, b)
        }
        fmt.Println("Result:", string(p.text))
    }
}
