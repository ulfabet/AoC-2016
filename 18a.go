package main

import "fmt"
import "crypto/md5"

//const input = ".^^.^.^^^^"
const input = ".^^^^^.^^.^^^.^...^..^^.^.^..^^^^^^^^^^..^...^^.^..^^^^..^^^^...^.^.^^^^^^^^....^..^^^^^^.^^^.^^^.^^"

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

var safe int

func main() {
    //walk("", 0, 0)
    //fmt.Println("Longest route:", length)
    //row := make([]byte, len(input))
    row := []byte(input)
    for j := 0; j < 40; j++ {
        fmt.Println(string(row))
        padded_input := "."+string(row)+"."
        for i := range row {
            if row[i] == '.' {
                safe++
            }
            if padded_input[i] != padded_input[i+2] { // it's a tarp!
                row[i] = '^'
            } else {
                row[i] = '.'
            }
        }
    }
    fmt.Println(safe)
}

