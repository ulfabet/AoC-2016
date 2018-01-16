package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

var positions []int
var start []int

func solve() int {
    for time := 0; time < 4000000; time++ {
        ok := true
        for i := range positions {
            //fmt.Println(time, i, (start[i]+time+i+1) % positions[i])
            ok = ok && (start[i]+time+i+1) % positions[i] == 0
        }
        if ok { return time }
    }
    return 0
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    n := 0
    for s.Scan() {
        var i int
        fields := strings.Fields(s.Text())
        fmt.Sscan(fields[3], &i)
        positions = append(positions, i)
        fmt.Sscan(fields[11], &i)
        start = append(start, i)
        n++
    }
    //for i := range positions { fmt.Println(positions[i], start[i]) }
    fmt.Println(solve())
}

