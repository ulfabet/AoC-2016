package main

import "fmt"
import "crypto/md5"

//const input = "hijkl"
const input = "ihgpwlah"
//const input = "kglvqrro"
//const input = "ulqzkmiv"
//const input = "gdjjyniy"

func modified_dragon_curve(a string) string {
    n := len(a)
    b := make([]byte, n)
    for i, v := range a {
        switch v {
            case '0': b[n-i-1] = '1'
            case '1': b[n-i-1] = '0'
        }
    }
    return a+"0"+string(b)
}

func checksum(s string) string {
    c := make([]byte, len(s)>>1)
    for i := range c {
        if s[2*i] == s[2*i+1] {
            c[i] = '1'
        } else {
            c[i] = '0'
        }
    }
    return string(c)
}

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
    //fmt.Println("x y =", x, y)
    hex := fmt.Sprintf("%x", md5.Sum([]byte(input+s)))
    //fmt.Println(s)
    //fmt.Println(hex)
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

func run() {
    //for tmp := 0; tmp < 4; tmp++ {
    for len(current) > 0 {
        fmt.Println(len(current))
        for _, v := range current {
            x, y := get_coordinates(v)
            if x == 3 && y == 3 {
                fmt.Println("Success!")
                fmt.Println(v)
                return
            }
            for _, w := range next_moves(v) {
                next = append(next, v+w)
            }
        }
        current, next = next, []string{}
    }
}

func main() {
    current = []string{""}
    run()
}

