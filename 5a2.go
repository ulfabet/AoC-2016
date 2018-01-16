package main

import (
    "fmt"
    "crypto/md5"
//    "log"
)

func main() {
    n := 3231929
    id := "abc"
    var s, x string
    var password string

    n = 0
    id = "reyedfim"

    var sum [16]byte

    fmt.Println("Working...")
    for {
        s = fmt.Sprintf("%s%d", id, n)
        sum = md5.Sum([]byte(s))
        if sum[0] == 0 && sum[1] == 0 {
            x = fmt.Sprintf("%X", sum)
            if x[:5] == "00000" {
                fmt.Println("Yes:", x[5:6])
                password += x[5:6]
                if len(password) == 8 { 
                    fmt.Println(password)
                    break
                }
            }
        }
        n++
        if n % 100000 == 0 { fmt.Println(n) }
    }
}
