package main

import (
    "fmt"
    "os"
    "bufio"
//    "log"
    "strings"
    "sort"
)

func process_line(line string) int {
    var id int
    var check string

    l := strings.Split(line, "-")

    m := make(map[int]string)
    a := strings.Join(l[:len(l)-1], "")
    for _, v := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
        m[strings.Count(a, v)] += v
    }
    //fmt.Println(m)

    i := 0
    keys := make([]int, len(m))
    for k := range m { keys[i] = k; i++ }

    sort.Sort(sort.Reverse(sort.IntSlice(keys)))

    result := ""
    for _, v := range keys {
        result += m[v]
    }
    //fmt.Println("result:", result[:5])


    b := l[len(l)-1]
    fmt.Sscanf(b, "%d[%5s]", &id, &check)
    //fmt.Println("> id:", id, "check:", check)

    if check != result[:5] { id = 0 }
    return id
}

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func main() {
    var cols [8]string
    var count [8]int
    var message [8]string


    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        line := s.Text()
        for i := 0; i < 8; i++ {
            cols[i] += string(line[i])
        }
    }

    for i := 0; i < 8; i++ {
        for _, v := range strings.Split(alphabet, "") {
            n := strings.Count(cols[i], v)
            if n > count[i] {
                count[i] = n
                message[i] = v
            }
        }
    }
    fmt.Println(strings.Join(message[:], ""))
}

