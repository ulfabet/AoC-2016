package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

type Screen struct {
    height int
    width int
    m [][]string
}

func atoi(s string) (i int) {
    fmt.Sscanf(s, "%d", &i)
    return
}

func (s *Screen) init() {
    s.height = 6
    s.width = 50
    s.m = make([][]string, s.height)
    for i := range s.m {
        s.m[i] = strings.Split(strings.Repeat(".", s.width), "")
    }
}
func (s *Screen) show() {
    for _, v := range s.m {
        fmt.Println(strings.Join(v, ""))
    }
}
func (s *Screen) rect(a, b int) {
    for i := 0; i < b; i++ {
        for j := 0; j < a; j++ {
            s.m[i][j] = "#"
        }
    }
}
func (s *Screen) rotate_row(a, b int) {
    for i := 0; i < b; i++ {
        tmp := s.m[a][s.width-1]
        for j := s.width-1; j > 0; j-- {
            s.m[a][j] = s.m[a][j-1]
        }
        s.m[a][0] = tmp
    }
}
func (s *Screen) rotate_column(a, b int) {
    for i := 0; i < b; i++ {
        tmp := s.m[s.height-1][a]
        for j := s.height-1; j > 0; j-- {
            s.m[j][a] = s.m[j-1][a]
        }
        s.m[0][a] = tmp
    }
}
func (s *Screen) count() (n int) {
    for i := range s.m {
        for j := range s.m[i] {
            if s.m[i][j] == "#" { n++ }
        }
    }
    return
}

func main() {
    var screen Screen

    /*
    screen.init()
    screen.rect(2, 3)
    screen.rotate_row(1, 49)
    screen.rotate_column(1, 4)
    screen.show()
    return
    */

    screen.init()
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        line := s.Text()
        fields := strings.Fields(line)
        switch fields[0] {
            case "rect": {
                xy := strings.Split(fields[1], "x")
                x, y := atoi(xy[0]), atoi(xy[1])
                screen.rect(x, y)
            }
            case "rotate": {
                a := atoi(strings.Split(fields[2], "=")[1])
                b := atoi(fields[4])
                switch fields[1] {
                    case "row": { screen.rotate_row(a, b) }
                    case "column": { screen.rotate_column(a, b) }
                }
            }
        }
    }
    screen.show()
    fmt.Println(screen.count())
}

