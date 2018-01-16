package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

type Set map[string]bool

type State struct{
    elevator int
    things [4]Set
    solution []string
    steps int
    finished bool
}

func (s *State) Init() {
    for i := range s.things {
        s.things[i] = make(Set)
    }
}

func (s *State) Add(floor int, thing string) {
    s.things[floor][thing] = true
}

func (s *State) Remove(floor int, thing string) {
    delete(s.things[floor], thing)
}

func (s *State) CurrentFloor() Set {
    return s.things[s.elevator]
}

func (s *State) Revert(direction int, things []string) {
    for _, thing := range things {
        s.Remove(s.elevator, thing)
    }
    s.elevator -= direction
    for _, thing := range things {
        s.Add(s.elevator, thing)
    }
    s.steps--
    s.solution = s.solution[:len(s.solution)-1]
}

func (s *State) Move(direction int, things []string) bool {
    if direction == 0 { return true }
    if direction == 1 && s.elevator == 3 { return false }
    if direction == -1 && s.elevator == 0 { return false }
    for _, thing := range things {
        s.Remove(s.elevator, thing)
    }
    s.elevator += direction
    for _, thing := range things {
        s.Add(s.elevator, thing)
    }
    foo := "down "
    if direction == 1 {
        foo = "up "
    }
    for _, thing := range things {
        foo += thing
    }
    s.solution = append(s.solution, foo)
    s.steps++
    return true
}

func (s *State) Illegal() bool {
    for _, v := range s.things {
        for thing := range v {
            if is_microchip(thing) && contains_generators(v) && !contains(v, compatible_generator(thing)) {
                //fmt.Println("Illegal:", s.things)
                return true
            }
        }
    }
    i := len(s.solution)
    if i > 1 {
        a := strings.Fields(s.solution[i-1])
        b := strings.Fields(s.solution[i-2])
        if a[1] == b[1] && a[0] != b[0] {
            //fmt.Println("Illegal: repeated move")
            return true
        }
    }
    return false
}

func (s *State) Finished() bool {
    if s.finished { return true }
    for i := 0; i < 3; i++ {
        if len(s.things[i]) > 0 { return false }
    }
    fmt.Println("Finished in", s.steps, "steps")
    fmt.Println(s.solution)
    s.finished = true
    return true
}

//----
func is_generator(thing string) bool {
    return thing[1] == 'g'
}

func is_microchip(thing string) bool {
    return thing[1] == 'm'
}

func compatible_generator(thing string) string {
    return fmt.Sprintf("%cg", thing[0])
}
func compatible_microchip(thing string) string {
    return fmt.Sprintf("%cm", thing[0])
}

func contains(s Set, thing string) bool {
    _, ok := s[thing]
    return ok
}

func contains_generators(s Set) bool {
    for k := range s {
        if k[1] == 'g' { return true }
    }
    return false
}

/*
func old_possible_selections(s Set) Set {
    r := make(Set)
    for t1 := range s {
        r.Add(t1)
        for t2 := range s {
            if t1 == t2 { continue }
            if contains(r, t2+t1) { continue }
            r.Add(t1+t2)
        }
    }
    return r
}
*/

func possible_selections(s Set) [][]string {
    var r [][]string
    var keys []string
    for k := range s {
        keys = append(keys, k)
    }
    for i, v := range keys {
        //r = append(r, []string{v})
        for j, w := range keys {
            if j <= i { continue }
            r = append(r, []string{v, w})
        }
    }
    for _, v := range keys {
        r = append(r, []string{v})
    }
    return r
}

func run(state State, level int) {
    if state.Finished() {
        return
    }
    if level > 20 {
        return
    }
    //fmt.Println(level)
    combinations := possible_selections(state.CurrentFloor())
    for _, things := range combinations {
        if !state.Move(1, things) {
            continue
        }
        if !state.Illegal() {
            run(state, level+1)
            if state.Finished() { return }
        }
        state.Revert(1, things)
    }
    for _, things := range combinations {
        if len(things) == 2 { continue }
        if !state.Move(-1, things) {
            continue
        }
        if !state.Illegal() {
            run(state, level+1)
            if state.Finished() { return }
        }
        state.Revert(-1, things)
    }
}

/*
func old_move(state State, direction int, things []string) {
    //fmt.Println("move", direction, things, "from floor", state.elevator)
    if !state.Move(direction, things) {
        return
    }
    if state.Illegal() {
        state.Revert(direction, things)
        return
    }
    //fmt.Println("ok")
    if state.Finished() {
        return
    }
    for _, v := range possible_selections(state.CurrentFloor()) {
        move(state, 1, v)
        if len(v) == 2 { continue }
        move(state, -1, v)
    }
}
*/


func main() {
    var state State
    state.Init()
    s := bufio.NewScanner(os.Stdin)
    floor := 0
    for s.Scan() {
        line := s.Text()
        fields := strings.Fields(line[:len(line)-1])
        for i := 0; i < len(fields); i++ {
            switch fields[i] {
            case "generator":
                fallthrough
            case "generator,":
                fallthrough
            case "microchip":
                fallthrough
            case "microchip,":
                thing := fmt.Sprintf("%c%c", fields[i-1][0], fields[i][0])
                state.Add(floor, thing)
            }
        }
        floor++
    }
    //move(state, 0, []string{})
    run(state, 0)
}

