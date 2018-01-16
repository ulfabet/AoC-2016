package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    var message [8]string
    var list [8][26]int

    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        for i, v := range s.Text() {
            list[i][int(v)-int('a')]++
        }
    }

    for i := range message {
        min := list[i][0]
        message[i] = "a"
        for k, v := range list[i] {
            if v < min {
                min = v
                message[i] = string(int('a')+k)
            }
        }
    }
    fmt.Println(strings.Join(message[:], ""))
}

