package main

import "fmt"
import "crypto/md5"

const input = "gdjjyniy"

var doors = []string{"U","D","L","R"}
var dx = []int{0,0,-1,1}
var dy = []int{-1,1,0,0}
var length int

func walk(route string, x, y int) {
    if x < 0 || x > 3 { return }
    if y < 0 || y > 3 { return }
    if x == 3 && y == 3 {
        if len(route) > length { length = len(route) }
        return
    }
    hex := fmt.Sprintf("%x", md5.Sum([]byte(input+route)))
    for i := range doors {
        if hex[i] > 97 { // door is open
            walk(route+doors[i], x+dx[i], y+dy[i])
        }
    }
}

func main() {
    walk("", 0, 0)
    fmt.Println("Longest route:", length)
}

