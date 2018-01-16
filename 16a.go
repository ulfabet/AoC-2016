package main

import "fmt"
//import "strings"

const input = "10011111011011001"
const disk_length = 272

//const input = "10000"
//const disk_length = 20

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
    var sum string
    for i := 0; i < len(s)-1; i += 2 {
        if s[i] == s[i+1] {
            sum = sum + "1"
        } else {
            sum = sum + "0"
        }
    }
    return sum
}

func main() {
    s := input
    for len(s) < disk_length {
        s = modified_dragon_curve(s)
    }
    fmt.Println(len(s), s)
    sum := checksum(s[:disk_length])
    for len(sum)&1 == 0 {
        sum = checksum(sum)
    }
    fmt.Println(len(sum), sum)
}

