package main

import (
    "fmt"
    "crypto/md5"
//    "log"
)

//const salt string = "abc"
const salt string = "yjdafjpo"

func get_hash(i int) string {
    return mem[i]
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

var mem [30000]string

func fill_mem() {
    for i := range mem {
        if i % 1000 == 0 { fmt.Println(i) }
        sum := md5.Sum([]byte(fmt.Sprintf("%s%d", salt, i)))
        for j := 0; j < 2016; j++ {
            //sum = md5.Sum(sum[:])
            sum = md5.Sum([]byte(fmt.Sprintf("%x", sum)))
        }
        mem[i] = fmt.Sprintf("%x", sum)
    }
}

func fill_mem_n(i int) {
    sum := md5.Sum([]byte(fmt.Sprintf("%s%d", salt, i)))
    for j := 0; j < 2016; j++ {
        //sum = md5.Sum(sum[:])
        sum = md5.Sum([]byte(fmt.Sprintf("%x", sum)))
    }
    mem[i] = fmt.Sprintf("%x", sum)
}

func mytest(i int) {
    fill_mem_n(i)
    hash := get_hash(i)
    fmt.Println(i, hash, get_repeat(hash))
}

func main() {
    fill_mem()
    index := 0
    number := 0
    for number < 64 {
        hash := get_hash(index)
        index++
        c := get_repeat(hash)
        if c == "" { continue }
        found := false
        for i := 0; i < 1000; i++ { // should be 1000
            if check_repeat(get_hash(i+index), c, 5) {
                found = true
                break
            }
        }
        if found {
            fmt.Println(index-1, hash)
            number++
        }
    }
}
