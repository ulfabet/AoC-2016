package main

import (
    "fmt"
    "os"
    "bufio"
)

var maze [][]int
var poi [8][]int
var maxx, maxy int

const space int = 0
const wall int = -1

var startx, starty int

func generate_maze() {
    fmt.Println("generate maze")
    s := bufio.NewScanner(os.Stdin)
    var y int
    for s.Scan() {
        line := s.Text()
        maxx = len(line)
        maze = append(maze, make([]int, len(line)))
        for x,c := range line {
            switch c {
                case '#': maze[y][x] = wall
                case '.': maze[y][x] = space
                case '0':
                    startx = x
                    starty = y
                    fallthrough
                default: // point of interest
                maze[y][x] = 0
                n := int(c-'0')
                poi[n] = []int{x, y}
                //fmt.Println("poi", n, x, y)
            }
        }
        maxy = y
        y++
    }
}

func show() {
    //fmt.Println("show")
    fmt.Printf("\x1b[2J\x1b[1;1H")
    for y, row := range maze {
        r := make([]byte, len(row))
        for x := range row {
            switch maze[y][x] {
            case wall:
                //fmt.Print("#")
                r[x] = '#'
            case space:
                r[x] = '.'
                /*
                for i,v := range poi {
                    if x == v[0] && y == v[1] {
                        c = fmt.Sprintf("%d", i)
                    }
                }
                fmt.Print(c)
                */
            default:
                r[x] = 'o'
                //fmt.Print("o")
            }
        }
        fmt.Println(string(r))
    }
}

func clear() {
    for y,row := range maze {
        for x := range row {
            if maze[y][x] > 0 {
                maze[y][x] = 0
            }
        }
    }
    fifo = [][]int{}
    fi = -1
}

var fifo [][]int
var fi int

func fifo_add(x, y, n int) {
    fifo = append(fifo, []int{x, y, n})
}

var total int
var best int

var prevn int

func walk(x, y, n int) {
    //fmt.Println("walk", x, y, n)
    if x < 0 || x > maxx { return }
    if y < 0 || y > maxy { return }
    if maze[y][x] == wall { return }
    if maze[y][x] != space && maze[y][x] <= n { return } // better path exists
    //old := maze[y][x]
    maze[y][x] = n
    ti := order[oi]
    if x == poi[ti][0] && y == poi[ti][1] {
        fmt.Println("Reached POI", ti, "in", n-1, "steps")
        total += n-1
        if oi == len(order)-1 {
            //fmt.Println("That's all folks!")
            if best == 0 || best > total { best = total }
            clear()
            return
        } else {
            clear()
            n = 1
            maze[y][x] = n
            oi++
        }
    }
    //fmt.Println("ok", x, y, n, old)
    if n > prevn {
        prevn = n
    //    show()
    }
    fifo_add(x-1, y, n+1)
    fifo_add(x, y-1, n+1)
    fifo_add(x+1, y, n+1)
    fifo_add(x, y+1, n+1)
}

var oi int
var order []int

var permutations [][]int
var pi int

func permgen(a []int, n int) {
    if n == 0 {
        b := make([]int, len(a))
        copy(b, a)
        permutations = append(permutations, b)
    } else {
        for i:=0; i<n; i++ {
            a[n-1], a[i] = a[i], a[n-1]
            permgen(a, n-1)
            a[n-1], a[i] = a[i], a[n-1]
        }
    }
}

func main() {
    generate_maze()

    list := []int{1,2,3,4,5,6,7}
    permgen(list, len(list))

    //permutations = append(permutations, []int{1,2,4,3})
    //permutations = append(permutations, []int{4,1,2,3})
    //permutations = append(permutations, []int{1,2,3,4})

    for pi=0; pi<len(permutations); pi++ {
        order = permutations[pi]
        oi = 0
        fmt.Println(order)
        fifo = [][]int{{startx, starty, 1}}
        fi = 0
        total = 0
        for fi < len(fifo) {
            //show()
            walk(fifo[fi][0], fifo[fi][1], fifo[fi][2])
            fi++
        }
        fmt.Println("Total:", total)
    }
    fmt.Println("Best:", best)
}

