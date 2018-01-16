package main

import "fmt"

func main() {
    count := []int{ 13, 19, 3, 7, 5, 17, 11 }
    start := []int{ 1, 10, 2, 1, 3, 5, 0 }
    for time := 0 ;; time++ {
        n := 0
        for i := range count {
            n += (start[i]+time+i+1) % count[i]
        }
        if n == 0 {
            fmt.Println(time)
            break
        }
    }
}

