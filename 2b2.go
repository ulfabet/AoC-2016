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
    x := 0
    y := 2
    m := [][]string{
    []string{"0", "0", "1", "0", "0"},
    []string{"0", "2", "3", "4", "0"},
    []string{"5", "6", "7", "8", "9"},
    []string{"0", "A", "B", "C", "0"},
    []string{"0", "0", "D", "0", "0"},
    }

    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        var line = s.Text()
        for _,v := range strings.Split(line, "") {
            oldx, oldy := x, y
            switch v {
                case "L": { x = max(0, x-1) }
                case "R": { x = min(4, x+1) }
                case "U": { y = max(0, y-1) }
                case "D": { y = min(4, y+1) }
                default: {
                    log.Println("Error")
                    break
                }
            }
            if m[y][x] == "0" {
                x, y = oldx, oldy
            }
        }
        log.Println("x", x, "y", y, m[y][x])
    }
}

