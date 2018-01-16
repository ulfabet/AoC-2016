package main

import (
//    "fmt"
    "os"
    "bufio"
    "log"
    "strings"
)

func min(i, j int) int {
    if i < j { return i }
    return j
}

func max(i, j int) int {
    if i > j { return i }
    return j
}

func main() {
    x := 2
    y := 2
    m := [][]int{
    []int{1, 2, 3},
    []int{4, 5, 6},
    []int{7, 8, 9},
    }

    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        var line = s.Text()
        //log.Println("line", line)
        for _,v := range strings.Split(line, "") {
            //log.Println(v)
            switch v {
                case "L": { x = max(1, x-1) }
                case "R": { x = min(3, x+1) }
                case "U": { y = max(1, y-1) }
                case "D": { y = min(3, y+1) }
                default: {
                    log.Println("Error")
                    break
                }
            }
        }
        log.Println("x", x, "y", y, m[y-1][x-1])
    }
}

