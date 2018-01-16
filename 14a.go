package main

import (
    "fmt"
    "crypto/md5"
//    "log"
)

//const salt string = "abc"
const salt string = "yjdafjpo"

func get_hash(s string, i int) string {
    sum := md5.Sum([]byte(fmt.Sprintf("%s%d", s, i)))
    return fmt.Sprintf("%x", sum)
}

func get_repeat(hash string) string {
    var prev rune
    n := 1
    for _,v := range hash {
        if v == prev { n++ } else { n = 1 }
        prev = v
        if n == 3 {
            return string(v)
        }
    }
    return ""
}

func check_repeat(hash string, c string, count int) bool {
    n := 0
    for _,v := range hash {
        if string(v) == c { n++ } else { n = 0 }
        if n == 5 { return true }
    }
    return false
}

func main() {
    index := 0
    number := 0
    for number < 64 {
        hash := get_hash(salt, index)
        index++
        //fmt.Println(hash)
        c := get_repeat(hash)
        if c == "" { continue }
        //fmt.Println(index-1, c)

        found := false
        for i := 0; i < 1000; i++ {
            if check_repeat(get_hash(salt, i+index), c, 5) {
                found = true
                break
            }
        }
        if found {
            fmt.Println(index-1, hash)
            number++
        } else {
            //fmt.Println("index", index-1, "does not produce a key")
        }
    }
}

/*
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
*/
