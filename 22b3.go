package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func min(a, b uint) uint {
    if a < b { return a }
    return b
}

func max(a, b uint) uint {
    if a > b { return a }
    return b
}

type Node struct {
    size int
    used int
    avail int
    goal bool
}

//const maxx = 2
//const maxy = 2
const maxx = 38
const maxy = 24

var nodes [maxy+1][maxx+1]*Node
var matrix [maxy+1][maxx+1]int
var gx,gy int
var total int

func add_node(n *Node, x, y int) {
    nodes[y][x] = n
    if y == 0 && x == maxx {
        n.goal = true
        gx, gy = x, y
        //tx, ty = x-1, y
    }
}

func show() {
    for y,row := range nodes {
        for x,n := range row {
            c := "."
            if n.used == 0 { c = "_" }
            if n.goal { c = "G" }
            if x == 0 && y == 0 {
                fmt.Printf("(%s)", c)
            } else {
                fmt.Printf(" %s ", c)
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

var fifo [][]int
var trail []string
var fi int

func fifo_add(x, y, n int, s string) {
    fifo = append(fifo, []int{x,y,n})
    trail = append(trail, s)
}

func fifo_clear() {
    fifo = [][]int{}
    trail = []string{}
    fi = -1
}

func is_empty(x, y int) bool {
    return nodes[y][x].used == 0
}

func is_goal(x, y int) bool {
    return nodes[y][x].goal
}

func clear_matrix() {
    for y, row := range matrix {
        for x := range row  {
            matrix[y][x] = 0
        }
    }
}

func follow_trail(x, y int, t string) {
    //fmt.Println("follow trail", x, y, t)
    if len(t) == 0 { return }
    var x2, y2 int
    switch t[len(t)-1] {
    case 'D': x2, y2 = x, y-1
    case 'U': x2, y2 = x, y+1
    case 'R': x2, y2 = x-1, y
    case 'L': x2, y2 = x+1, y
    }
    move(x2,y2,x,y)
    //show()
    follow_trail(x2,y2,t[:len(t)-1])
}

func try_to_empty() {
    x,y,n := fifo[fi][0], fifo[fi][1], fifo[fi][2]
    t := trail[fi]
    //fmt.Println("try to empty", x, y, n)
    if x < 0 || y < 0 { return }
    if x > maxx || y > maxy { return }
    if matrix[y][x] != 0 && n >= matrix[y][x] { return }
    if is_goal(x, y) { return }
    //show()
    if is_empty(x, y) {
        fmt.Println("Steps", n, t)
        total += n
        follow_trail(x, y, t)
        fifo_clear()
        clear_matrix()
        return
    }
    matrix[y][x] = n
    if valid_move(x,y,x+1,y) { fifo_add(x+1, y, n+1, t+"R") }
    if valid_move(x,y,x,y+1) { fifo_add(x, y+1, n+1, t+"D") }
    if valid_move(x,y,x-1,y) { fifo_add(x-1, y, n+1, t+"L") }
    if valid_move(x,y,x,y-1) { fifo_add(x, y-1, n+1, t+"U") }
}

func try_move_left(x, y int) {
    //fmt.Println("try move left", x, y)
    if x == 0 && y == 0 {
        fmt.Println("Finished in", total, "steps")
        return
    }
    if !valid_move(x, y, x-1, y) {
        fmt.Println("Oops!")
        return
    }
    fifo_add(x-1, y, 1, "L")
    for fi < len(fifo) {
        try_to_empty()
        fi++
    }
    try_move_left(gx, gy)
}

func valid_move(x, y, x2, y2 int) bool {
    if x < 0 || x > maxx { return false }
    if x2 < 0 || x2 > maxx { return false }
    if y < 0 || y > maxy { return false }
    if y2 < 0 || y2 > maxy { return false }
    s := nodes[y][x]
    d := nodes[y2][x2]
    return d.size >= s.used
}

func move(x, y, x2, y2 int) {
    s := nodes[y][x]
    d := nodes[y2][x2]
    d.used += s.used
    d.avail -= s.used
    d.goal = s.goal
    s.used = 0
    s.avail = s.size
    s.goal = false
    if d.goal {
        gx, gy = x-1, y
    }
    if d.avail < 0 {
        fmt.Println("Error", d)
    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    s.Scan()
    for s.Scan() {
        var x,y,a,b,c int
        f := strings.Fields(s.Text())
        fmt.Sscanf(f[0], "/dev/grid/node-x%d-y%d", &x, &y)
        fmt.Sscan(f[1], &a)
        fmt.Sscan(f[2], &b)
        fmt.Sscan(f[3], &c)
        n := Node{a, b, c, false}
        add_node(&n, x, y)
    }
    try_move_left(gx,gy)
}
