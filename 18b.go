package main

import "fmt"

//const input = ".^^.^.^^^^"
const input = ".^^^^^.^^.^^^.^...^..^^.^.^..^^^^^^^^^^..^...^^.^..^^^^..^^^^...^.^.^^^^^^^^....^..^^^^^^.^^^.^^^.^^"

var safe int

func main() {
    row := []byte(input)
    for j := 0; j < 400000; j++ {
        //fmt.Println(string(row))
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

