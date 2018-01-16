package main

import "fmt"

const input = "10011111011011001"
//const disk_length = 272
const disk_length = 35651584

func modified_dragon_curve(a string) string {
    n := len(a)
    b := make([]byte, n)
    for i, v := range a {
        switch v {
            case '0': b[n-i-1] = '1'
            case '1': b[n-i-1] = '0'
        }
    }
    return a+"0"+string(b)
}

func checksum(s string) string {
    c := make([]byte, len(s)>>1)
    for i := range c {
        if s[2*i] == s[2*i+1] {
            c[i] = '1'
        } else {
            c[i] = '0'
        }
    }
    return string(c)
}

func main() {
    s := input
    for len(s) < disk_length {
        s = modified_dragon_curve(s)
    }
    //fmt.Println(len(s), s)

    sum := checksum(s[:disk_length])
    for len(sum)&1 == 0 {
        sum = checksum(sum)
    }
    fmt.Println(len(sum), sum)
}

