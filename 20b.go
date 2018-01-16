package main

import (
    "fmt"
    "os"
    "bufio"
//    "strings"
)

func min(a, b uint) uint {
    if a < b { return a }
    return b
}

func max(a, b uint) uint {
    if a > b { return a }
    return b
}

type Interval struct {
    low uint
    high uint
}

func (a *Interval) contains(b uint) bool {
    return a.low <= b && b <= a.high
}

func (a *Interval) overlaps(b *Interval) bool {
    return a.low <= b.high && a.high >= b.low
}

func (a *Interval) join(b *Interval) *Interval {
    return &Interval{min(a.low, b.low), max(a.high, b.high)}
}

type Set map[*Interval]bool

func (a Set) add(b *Interval) {
    a[b] = true
}

func (a Set) remove(b *Interval) {
    delete(a, b)
}

func join_overlapping_intervals(set Set) bool {
    // todo pairs
    for a := range set {
        //fmt.Println("check", a)
        for b := range set {
            if a == b { continue }
            if a.overlaps(b) {
                set.remove(a)
                set.remove(b)
                set.add(a.join(b))
                return true
            }
        }
    }
    return false
}

func main() {
    set := make(Set)
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        var i, j uint
        fmt.Sscanf(s.Text(), "%d-%d", &i, &j)
        set.add(&Interval{i, j})
    }
    for join_overlapping_intervals(set) {}

    var n uint = 0xffffffff
    for a := range set {
        n -= a.high-a.low+1
    }
    fmt.Println(n+1)
}
