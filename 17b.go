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
    x, y := get_coordinates(s)
    hex := fmt.Sprintf("%x", md5.Sum([]byte(input+s)))
    for i := range directions {
        if hex[i] > 97 { // door is open
            dir := directions[i]
            switch {
                case dir == "U" && y == 0: continue
                case dir == "D" && y == 3: continue
                case dir == "L" && x == 0: continue
                case dir == "R" && x == 3: continue
            }
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

