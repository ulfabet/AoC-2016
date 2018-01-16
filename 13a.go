package main

import (
    "fmt"
)

func atoi(s string) (i int) {
    fmt.Sscan(s, &i)
    return
}

/*
const maxx int = 9
const maxy int = 6
var maze [7][10]int
const favourite_number int = 10
var target []int = []int{7, 4}
*/

const maxx int = 50
const maxy int = 50
var maze [maxy+1][maxx+1]int
const favourite_number int = 1362
var target []int = []int{31, 39}

var steps int
const space int = 0
const wall int = 1


func generate_maze() {
    for x := 0; x <= maxx; x++ {
        for y := 0; y <= maxy; y++ {
            f := x*x+3*x+2*x*y+y+y*y
            f += favourite_number
            n := 0
            for f > 0 {
                n ^= f&1
                f = f>>1
            }
            maze[y][x] = -n
        }
    }
}

func show() {
    for y, row := range maze {
        for x := range row {
            switch maze[y][x] {
            case -1:
                fmt.Print("#")
            case 0:
                fmt.Print(".")
            default:
                fmt.Print("O")
            }
        }
        fmt.Println()
    }
}

func generate_tree(x, y, n int) {
    if x < 0 || y < 0 { return }
    if x > maxx || y > maxy { return }
    if maze[y][x] == wall { return }
    if maze[y][x] != space && maze[y][x] < n { return } // better path exists
    maze[y][x] = n
    if x == target[0] && y == target[1] {
        fmt.Println("Finished in", n-1, "steps")
        return
    }
    generate_tree(x-1, y, n+1)
    generate_tree(x, y-1, n+1)
    generate_tree(x+1, y, n+1)
    generate_tree(x, y+1, n+1)
}

func main() {
    generate_maze()
    generate_tree(1, 1, 1)
    //show()
}

