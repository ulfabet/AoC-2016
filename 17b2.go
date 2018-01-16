package main

import "fmt"
import "crypto/md5"

//const input = "hijkl"
//const input = "ihgpwlah"
//const input = "kglvqrro"
//const input = "ulqzkmiv"
const input = "gdjjyniy"

var directions = []string{"U","D","L","R"}

func get_coordinates(s string) (x, y int) {
    for _, v := range s {
        switch v {
        case 'U': y--
        case 'D': y++
        case 'L': x--
        case 'R': x++
        }
    }
    return
}

func next_moves(s string) (moves []string) {
    hex := fmt.Sprintf("%x", md5.Sum([]byte(input+s)))
    for i := range directions {
        if hex[i] > 97 { // door is open
            moves = append(moves, directions[i])
        }
    }
    return
}

var current []string
var next []string
var success int

func run() {
    for len(current) > 0 {
        fmt.Println(len(current))
        for _, v := range current {
            x, y := get_coordinates(v)
            if x == 3 && y == 3 {
                //fmt.Println("Success!", len(v))
                success = len(v)
            } else {
                for _, w := range next_moves(v) {
                    switch {
                        case w == "U" && y == 0: continue
                        case w == "D" && y == 3: continue
                        case w == "L" && x == 0: continue
                        case w == "R" && x == 3: continue
                    }
                    next = append(next, v+w)
                }
            }
        }
        current, next = next, []string{}
    }
}

func main() {
    current = []string{""}
    run()
    fmt.Println("Success!", success)
}

