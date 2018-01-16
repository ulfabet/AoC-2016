package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

const input = "abcdefgh"
//const input = "abcde"

func min(a, b uint) uint {
    if a < b { return a }
    return b
}

func max(a, b uint) uint {
    if a > b { return a }
    return b
}

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
    n := len(p.text)
    a %= n
    switch direction {
    case "left":
        p.text = append(p.text[a%n:], p.text[:a%n]...)
    case "right":
        p.text = append(p.text[(n-a)%n:], p.text[:(n-a)%n]...)
    default:
        fmt.Println("rotate error", a)
    }
}

func (p *Password) rotate_based(a byte) {
    var i int
    for i = range p.text {
        if a == p.text[i] {
            if i >= 4 { i++ }
            p.rotate("right", i+1)
            return
        }
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

type Node struct {
    size int
    used int
    avail int
    percent int
}

var nodes []*Node
var pairs [][]*Node

func main() {
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    s.Scan()
    for s.Scan() {
        var a,b,c,d int
        f := strings.Fields(s.Text())
        fmt.Sscan(f[1], &a)
        fmt.Sscan(f[2], &b)
        fmt.Sscan(f[3], &c)
        fmt.Sscan(f[4], &d)
        //fmt.Println(f, a, b, c, d)
        n := Node{a, b, c, d}
        nodes = append(nodes, &n)
    }
    fmt.Println(len(nodes))
    for i,a := range nodes {
        for j,b := range nodes {
            if i == j { continue }
            if a.used == 0 { continue }
            if a.used > b.avail { continue }
            pairs = append(pairs, []*Node{a, b})
        }
    }
    fmt.Println(len(pairs))
}
