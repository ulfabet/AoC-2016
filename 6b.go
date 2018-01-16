package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func main() {
    var cols [8]string
    var count [8]int
    var message [8]string

    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        line := s.Text()
        for i := range message {
            cols[i] += string(line[i])
        }
    }

    for i := range message {
        for _, v := range strings.Split(alphabet, "") {
            n := strings.Count(cols[i], v)
            if count[i] == 0 || n < count[i] {
                count[i] = n
                message[i] = v
            }
        }
    }
    fmt.Println(strings.Join(message[:], ""))
}

