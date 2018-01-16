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
type Arg struct{
    x,y int
    xreg,yreg bool
}

type Processor struct {
    pc int
    reg [4]int
    code []func(Arg)
    args []Arg
}

func (p *Processor) init(a,b,c,d int) {
    p.pc = 0
    p.reg[0] = a
    p.reg[1] = b
    p.reg[2] = c
    p.reg[3] = d
}

func (p *Processor) op_cpy(a Arg) {
    if a.y < 0 || a.y > len(p.reg)-1 { return }
    if a.xreg {
        p.reg[a.y] = p.reg[a.x]
    } else {
        p.reg[a.y] = a.x
    }
}

func (p *Processor) op_inc(a Arg) {
    p.reg[a.x]++
}

func (p *Processor) op_dec(a Arg) {
    p.reg[a.x]--
}

func (p *Processor) op_jnz(a Arg) {
    var n int
    if a.yreg {
        n = p.reg[a.y]-1
    } else {
        n = a.y-1
    }
    if a.xreg {
        if p.reg[a.x] != 0 { p.pc += n }
    } else {
        if a.x != 0 { p.pc += n }
    }
}

func (p *Processor) op_out(a Arg) {
    if a.xreg {
        out[oi] = p.reg[a.x]
    } else {
        out[oi] = a.x
    }
    oi++
    if oi == len(out) { p.pc = len(p.code) } // exit
}

func (p *Processor) run() {
    for p.pc >= 0 && p.pc < len(p.code) {
        code := p.code[p.pc]
        args := p.args[p.pc]
        code(args)
        p.pc++
    }
}

func reg_or_value(a string) (int,bool) {
    switch a {
        case "a": fallthrough
        case "b": fallthrough
        case "c": fallthrough
        case "d":
            return int(a[0]-'a'), true
        default:
            return atoi(a), false
    }
}

func (p *Processor) add_code(line string) {
    var x,y int
    var xreg, yreg bool
    var code func(Arg)
    field := strings.Fields(line)
    switch field[0] {
    case "cpy":
        code = p.op_cpy
        x, xreg = reg_or_value(field[1])
        y, yreg = reg_or_value(field[2])
    case "inc":
        code = p.op_inc
        x, xreg = reg_or_value(field[1])
    case "dec":
        code = p.op_dec
        x, xreg = reg_or_value(field[1])
    case "jnz":
        code = p.op_jnz
        x, xreg = reg_or_value(field[1])
        y, yreg = reg_or_value(field[2])
    case "out":
        code = p.op_out
        x, xreg = reg_or_value(field[1])
    }
    p.code = append(p.code, code)
    p.args = append(p.args, Arg{x,y,xreg,yreg})
}

var out [8]int
var oi int

// --------
func main() {
    var p Processor
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        p.add_code(s.Text())
    }

    a := 0
    for out != [8]int{0,1,0,1,0,1,0,1} {
        fmt.Println("a", a)
        p.init(a,0,0,0)
        oi = 0
        p.run()
        fmt.Println(p.reg)
        fmt.Println(out)
        a++
    }
}

