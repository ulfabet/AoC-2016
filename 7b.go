package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func main() {

    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        line := s.Text()
        var outside, inside []string
        x := &outside
        for _, v := range strings.Split(line, "[") {
            for _, w := range strings.Split(v, "]") {
                *x = append(*x, w)
                switch x {
                    case &outside: { x = &inside }
                    case &inside: { x = &outside }
                }
            }
        }

        var s string
        var matchOut, matchIn bool

        for i, a := range alphabet {
            for j, b := range alphabet {
                if i == j { continue }
                s = fmt.Sprintf("%c%c%c", a, b, a)
                matchOut = strings.Contains(strings.Join(outside, " "), s)
                s = fmt.Sprintf("%c%c%c", b, a, b)
                matchIn = strings.Contains(strings.Join(inside, " "), s)

                /*
                s = fmt.Sprintf("%c%c%c", a, b, a)
                for _, v := range outside {
                    if strings.Contains(v, s) {
                        matchOut = true
                        break
                    }
                }
                s = fmt.Sprintf("%c%c%c", b, a, b)
                for _, v := range inside {
                    if strings.Contains(v, s) {
                        matchIn = true
                        break
                    }
                }
                */
                if matchOut && matchIn { break }
            }
            if matchOut && matchIn { break }
        }
        if matchOut && matchIn {
            fmt.Println("Match:", line)
        }
    }
}

