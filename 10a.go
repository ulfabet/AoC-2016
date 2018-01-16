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
type Receiver interface {
    give(n int)
}

// --------
type Output struct {
    number int
    value int
}

func (o *Output) give(n int) {
    //fmt.Println("give output", o.number, "value", n)
    o.value = n
}

// --------
type Bot struct {
    number int
    values []int
    low_to *Receiver
    high_to *Receiver
}

func (b *Bot) action() int {
    //fmt.Println("bot", b.number, "action, values:", b.values)
    if len(b.values) < 2 {
        return 0
    }
    low, high := b.values[0], b.values[1]
    b.values = []int{}
    if low > high { low, high = high, low }
    if low == 17 && high == 61 {
        fmt.Println("Bingo for bot", b.number)
    }
    (*b.low_to).give(low)
    (*b.high_to).give(high)
    return 1
}

/*
func (b *Bot) ready() bool {
    // should we also check whether receivers are not nil?
    return (len(b.values) == 2)
}
*/

func (b *Bot) give(n int) {
    //fmt.Println("give bot", b.number, "value", n)
    b.values = append(b.values, n)
    //if len(b.values) == 2 {
    //    b.action()
    //}
}

// --------
var bots map[int]*Bot
var outputs map[int]*Output

func get_bot(n int) *Bot {
    b, ok := bots[n]
    if !ok {
        b = new(Bot)
        b.number = n
        bots[n] = b
    }
    return b
}

func get_output(n int) *Output {
    b, ok := outputs[n]
    if !ok {
        b = new(Output)
        b.number = n
        outputs[n] = b
    }
    return b
}

func process_line(line string) {
    fields := strings.Fields(line)
    switch fields[0] {
        case "value": {
            b := get_bot(atoi(fields[5]))
            b.give(atoi(fields[1]))
        }
        case "bot": {
            var b, c Receiver
            a := get_bot(atoi(fields[1]))
            if fields[5] == "bot" {
                b = get_bot(atoi(fields[6]))
            } else {
                b = get_output(atoi(fields[6]))
            }
            if fields[10] == "bot" {
                c = get_bot(atoi(fields[11]))
            } else {
                c = get_output(atoi(fields[11]))
            }
            a.low_to = &b
            a.high_to = &c
        }
    }
}

func main() {
    bots = make(map[int]*Bot)
    outputs = make(map[int]*Output)
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        process_line(s.Text())
    }
    var actions int
    for {
        actions = 0
        for _, v := range bots {
            actions += v.action()
        }
        if actions == 0 { break }
    }
    for _, v := range outputs {
        fmt.Println("output", v.number, "value", v.value)
    }
    fmt.Println("part2:", outputs[0].value*outputs[1].value*outputs[2].value)
}

