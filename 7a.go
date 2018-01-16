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

        match := false
        illegal := false
        for i, a := range alphabet {
            for j, b := range alphabet {
                if i == j { continue }

                s := fmt.Sprintf("%c%c%c%c", a, b, b, a)
                for _, v := range inside {
                    if strings.Contains(v, s) {
                        illegal = true
                    }
                }
                for _, v := range outside {
                    if strings.Contains(v, s) {
                        match = true
                    }
                }
            }
        }
        if match && !illegal {
            fmt.Println("Match:", line)
        }
    }
}

