package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func get_aba_list(s string) [][]string {
    for i := 0; i < len(s)-3; i++ {
        var value [][]string
        a, b := s[i], s[i+1]
        if a == s[i+2] {
            value = append(value, {a, b})
        }
        return value
}

func has_aba(s string) bool {
    for i := 0; i < len(s)-3; i++ {
        a, b := s[i], s[i+1]
        if a != b && a == s[i+2] && has_bab(s, a, b) {
            return true
        }
    }
    return false
}

func has_bab(s, a, b string) bool {
    for i := 0; i < len(s)-3; i++ {
        if s[i] == b && s[i+1] == a && s[i+2] == b {
            return true
        }
    }
    return false
}

func main() {

    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        line := s.Text()
        for i, v := range get_aba_list(line) {
            a, b := v[0], v[1]
            if has_bab(line a, b) {
                fn.Println("Match", line)
            }
        }
    }
    return

    /*
    for {
        var t1, t2, t3 string
        _, err := fmt.Scanf("%s[%s]%s", &t1, &t2, &t3)
        if err != nil {
            fmt.Println("err", err)
            break
        }
        fmt.Println("test", t1, t2, t3)
    }
    return
    */

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

        out := strings.Join(outside, " ")
        in := strings.Join(inside, " ")

        var aba, bab bool
        for i := 0; i < len(out)-3; i++ {
            if out[i] == out[i+2] && out[i] != out[i+1] {
                aba = true
                if bab == true { return }
            }
        }


        for i, a := range alphabet {
            for j, b := range alphabet {
                if i == j { continue }
                s = fmt.Sprintf("%c%c%c", a, b, a)
                matchOut = strings.Contains(out, s)
                s = fmt.Sprintf("%c%c%c", b, a, b)
                matchIn = strings.Contains(in, s)
                if matchOut && matchIn { break }
            }
            if matchOut && matchIn { break }
        }
        if matchOut && matchIn {
            fmt.Println("Match:", line)
        }
    }
}

