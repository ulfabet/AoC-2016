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
    //percent int
    goal bool
    too_big bool
}

//const maxx = 2
//const maxy = 2
const maxx = 38
const maxy = 24

var nodes [maxy+1][maxx+1]*Node
var pairs [][]*Node

func add_node(n *Node, x, y int) {
    nodes[y][x] = n
    if y == 0 && x == maxx { n.goal = true }
}

func show() {
    for y,row := range nodes {
        for x,n := range row {
            c := "."
            if n.used == 0 { c = "_" }
            if n.goal { c = "G" }
            if n.too_big { c = "#" }
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

func neighbours(x, y int) []*Node {
    var r []*Node
    if x > 0 { r = append(r, nodes[y][x-1]) }
    if y > 0 { r = append(r, nodes[y-1][x]) }
    if x < maxx { r = append(r, nodes[y][x+1]) }
    if y < maxy { r = append(r, nodes[y+1][x]) }
    return r
}

func neighbours_coords(x, y int) [][]int {
    var r [][]int
    if x > 0 { r = append(r, []int{x-1,y}) }
    if y > 0 { r = append(r, []int{x,y-1}) }
    if x < maxx { r = append(r, []int{x+1,y}) }
    if y < maxy { r = append(r, []int{x,y+1}) }
    return r
}

func valid_moves(x, y int) [][]int {
    var r [][]int
    if x > 0 && valid_move(x,y,x-1,y) { r = append(r, []int{x-1,y}) }
    if y > 0 && valid_move(x,y,x,y-1) { r = append(r, []int{x,y-1}) }
    if x < maxx && valid_move(x,y,x+1,y) { r = append(r, []int{x+1,y}) }
    if y < maxy && valid_move(x,y,x,y+1) { r = append(r, []int{x,y+1}) }
    return r
}

func valid_move(x1,y1,x2,y2 int) bool {
    if last_move == [4]int{x2,y2,x1,y1} { return false } // trivial loop
    s := nodes[y1][x1]
    d := nodes[y2][x2]
    if s.goal && x2 > x1 { return false } // only move left?
    return s.used > 0 && d.avail > s.used // no use moving empty node
}

var last_move [4]int = [4]int{0,0,0,0}
var steps int

func move(x1,y1,x2,y2 int) int {
    last_move = [4]int{x1,y1,x2,y2}
    steps++
    s := nodes[y1][x1]
    d := nodes[y2][x2]
    r := s.used
    d.used += s.used
    d.avail -= s.used
    d.goal = s.goal
    s.used = 0
    s.avail = s.size
    s.goal = false
    return r
}

func revert(x1,y1,x2,y2,data int) {
    steps--
    s := nodes[y1][x1]
    d := nodes[y2][x2]
    s.used += data
    s.avail -= data
    s.goal = d.goal
    d.used -= data
    d.avail += data
    d.goal = false
}

var best int = 50

// todo: avoid loops?
func search_neighbours(x, y int) {
    //fmt.Println("search n", x, y, steps)
    //show()
    if nodes[0][0].goal {
        fmt.Println("Goal!", steps)
        best = steps
        return
    }
    if steps > best {
        return
    }
    for _,c := range neighbours_coords(x, y) { // todo: coords
        //fmt.Println(c)
        if !valid_move(c[0], c[1], x, y) { continue }
        //fmt.Println("move")
        data := move(c[0], c[1], x, y)
        search_neighbours(c[0], c[1])
        revert(c[0], c[1], x, y, data)
    }
    //fmt.Println("return")
}

func search() {
    //fmt.Println("search", steps)
    //show()
    for y,row := range nodes {
        for x := range row {
            for _,c := range valid_moves(x,y) {
                data := move(x, y, c[0], c[1])
                search_neighbours(x, y)
                revert(x, y, c[0], c[1], data)
            }
        }
    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)
    s.Scan()
    s.Scan()
    for s.Scan() {
        var a,b,c,d int
        var x,y int
        f := strings.Fields(s.Text())
        fmt.Sscanf(f[0], "/dev/grid/node-x%d-y%d", &x, &y)
        fmt.Sscan(f[1], &a)
        fmt.Sscan(f[2], &b)
        fmt.Sscan(f[3], &c)
        fmt.Sscan(f[4], &d)
        n := Node{a, b, c, false, false}
        add_node(&n, x, y)
    }
    for y,row := range nodes {
        for x,n := range row {
            //fmt.Println(n)
            too_big := true
            for _,m := range neighbours(x, y) {
                if n.used > m.size { continue }
                too_big = false
            }
            if too_big {
                //fmt.Println("found too big")
                n.too_big = true
            }
        }
    }

    fmt.Println("searching...")
    search()
    fmt.Println("best", best)

    //show()

    /*
    move(1,0,1,1)
    show()
    move(2,0,1,0)
    show()
    */

    /*
    for i,a := range nodes {
        for j,b := range nodes {
            if i == j { continue }
            if a.used == 0 { continue }
            if a.used > b.avail { continue }
            pairs = append(pairs, []*Node{a, b})
        }
    }
    fmt.Println(len(pairs))
    */
}
