package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
//    "log"
//    "sort"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func rotate(s string, n int) string {
    b := []byte(s)
    for i, v := range b {
        if v != byte(' ') {
            b[i] = alphabet[(int(b[i])-int(byte('a'))+n) % 26]
        }
    }
    return string(b)
}

func process_line(line string) {
    l := strings.Split(line, "-")
    name := strings.Join(l[:len(l)-1], " ")

    var id int
    var check string
    fmt.Sscanf(l[len(l)-1], "%d[%5s]", &id, &check)

    fmt.Println(rotate(name, id), id)
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        process_line(s.Text())
    }
}

