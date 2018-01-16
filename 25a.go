package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func atoi(s string) (i int) {
    fmt.Sscan(s, &i)
    return
}

// --------
type Processor struct {
    pc int
    reg [4]int
    code []string
}

func (p *Processor) init(a,b,c,d int) {
    p.pc = 0
    p.reg[0] = a
    p.reg[1] = b
    p.reg[2] = c
    p.reg[3] = d
}

func (p *Processor) toggle(pc int) {
    if pc >= 0 && pc < len(p.code) {
        line := p.code[pc]
        field := strings.Fields(line)
        if len(field) == 2 {
            switch field[0] {
            case "inc":
                field[0] = "dec"
                p.code[pc] = strings.Join(field, " ")
            default:
                field[0] = "inc"
                p.code[pc] = strings.Join(field, " ")
            }
            return
        }
        if len(field) == 3 {
            switch field[0] {
            case "jnz":
                field[0] = "cpy"
                p.code[pc] = strings.Join(field, " ")
            default:
                field[0] = "jnz"
                p.code[pc] = strings.Join(field, " ")
            }
            return
        }
    }
}

func (p *Processor) run() {
    for p.pc >= 0 && p.pc < len(p.code) {
        //fmt.Print(p.pc, ": ")
        line := p.code[p.pc]
        //fmt.Println(line, p.reg)
        field := strings.Fields(line)
        switch field[0] {
        case "cpy":
            var x, y int
            y = int(field[2][0]-'a')
            if y < 0 || y > len(p.reg)-1 { continue }
            //fmt.Println("y", y, field[2])
            switch field[1] {
            case "a": fallthrough
            case "b": fallthrough
            case "c": fallthrough
            case "d":
                x = int(field[1][0]-'a')
                //fmt.Println("x register", x, field[1])
                p.reg[y] = p.reg[x]
            default:
                x = atoi(field[1])
                //fmt.Println("x value", x, field[1])
                p.reg[y] = x
            }
        case "inc":
            x := int(field[1][0]-'a')
            p.reg[x]++
        case "dec":
            x := int(field[1][0]-'a')
            p.reg[x]--
        case "jnz":
            var x, y int
            y = int(field[2][0]-'a')
            if y < 0 || y > len(p.reg)-1 {
                y = atoi(field[2])
            } else {
                y = p.reg[y]
            }
            switch field[1] {
            case "a": fallthrough
            case "b": fallthrough
            case "c": fallthrough
            case "d":
                x = int(field[1][0]-'a')
                if p.reg[x] != 0 { p.pc += y-1 }
            default:
                x = atoi(field[1])
                if x != 0 { p.pc += y-1 }
            }
        case "tgl":
            var x int
            switch field[1] {
            case "a": fallthrough
            case "b": fallthrough
            case "c": fallthrough
            case "d":
                x = int(field[1][0]-'a')
                p.toggle(p.pc+p.reg[x])
            default:
                x = atoi(field[1])
                p.toggle(p.pc+x)
            }
        case "out":
            var x int
            switch field[1] {
            case "a": fallthrough
            case "b": fallthrough
            case "c": fallthrough
            case "d":
                x = int(field[1][0]-'a')
            default:
                x = atoi(field[1])
            }
            out[oi] = x
            oi++
            if oi == len(out) { p.pc = len(p.code) }
        }
        p.pc++
    }
}

var out [8]int
var oi int

// --------
func main() {
    var p Processor
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        p.code = append(p.code, s.Text())
    }

    a := 1200
    for out != [8]int{0,1,0,1,0,1,0,1} {
        fmt.Println("a", a)
        p.init(a,0,0,0)
        oi = 0
        p.run()
        fmt.Println(out)
        fmt.Println(p.reg)
        a += 1
    }
}

