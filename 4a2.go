package main

import (
    "fmt"
    "os"
    "bufio"
//    "log"
    "strings"
    "sort"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz"

func process_line(line string) int {
    var id int
    var check string

    l := strings.Split(line, "-")

    m := make(map[int]string)
    a := strings.Join(l[:len(l)-1], "")
    //for _, v := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
    for _, v := range alphabet {
        w := string(v)
        m[strings.Count(a, w)] += w
    }

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

func main() {
    var sum int

    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        sum += process_line(s.Text())
    }
    fmt.Println("sum:", sum)
}

