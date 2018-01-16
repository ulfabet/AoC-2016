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

    n = 0
    id = "reyedfim"

    var sum [16]byte
    var newpass []byte
    var pos int
    var ready bool

    newpass = []byte("________")

    newpass = []byte("___D_E_7")
    n = 3000000

    var data []byte
    //data = []byte("abc") + []byte("3231929")
    data = []byte("abc")
    data[3:] = []byte("3231929")
    fmt.Println(data)
    return

    fmt.Println("Working...")
    for ready == false {
        s = fmt.Sprintf("%s%d", id, n)
        sum = md5.Sum([]byte(s))
        if sum[0] == 0 && sum[1] == 0 {
            x = fmt.Sprintf("%X", sum)
            if x[:5] == "00000" {
                fmt.Println("Yes:", x[5:6])
                pos = int(x[5] - byte('0'))
                if pos < len(newpass) && newpass[pos] == byte('_') {
                    newpass[pos] = x[6]
                    fmt.Println(string(newpass))
                }
                ready = true
                for _, v := range newpass {
                    if v == byte('_') { ready = false }
                }
            }
        }
        n++
        if n % 100000 == 0 { fmt.Println(n) }
    }
    fmt.Println(string(newpass))
}
