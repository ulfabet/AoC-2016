package main

import (
    "fmt"
    "crypto/md5"
)

const salt string = "abc"
//const salt string = "yjdafjpo"

func old_get_hash(s string, i int) string {
    sum := md5.Sum([]byte(fmt.Sprintf("%s%d", s, i)))
    return fmt.Sprintf("%x", sum)
}

func old_get_repeat(hash string) string {
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

func old_check_repeat(hash string, c string, count int) bool {
    n := 0
    for _,v := range hash {
        if string(v) == c { n++ } else { n = 0 }
        if n == 5 { return true }
    }
    return false
}

var mem [30000][16]byte

func get_hash(i int) [16]byte {
    return mem[i]
}

func get_repeat(hash [16]byte) string {
    var prev rune
    n := 1
    for _,v := range fmt.Sprintf("%x", hash) {
        if v == prev { n++ } else { n = 1 }
        prev = v
        if n == 3 { return string(v) }
    }
    return ""
}

func check_repeat(hash [16]byte, c string, count int) bool {
    n := 0
    for _,v := range hash {
        if v == c { n++ } else { n = 0 }
        if n == 5 { return true }
    }
    return false
}

func setup() {
    for i := range mem {
        if i % 1000 == 0 { fmt.Println(i) }
        sum := md5.Sum([]byte(fmt.Sprintf("%s%d", salt, i)))
        mem[i] = sum
    }
}

func main() {
    setup()
    index := 18
    number := 0
    for number < 64 {
        hash := get_hash(index)
        fmt.Printf("%d, %x\n", index, hash)
        index++
        c := get_repeat(hash)
        fmt.Println("c", c)
        break
        if c == 0 { continue }
        //fmt.Println(index-1, c)

        found := false
        for i := 0; i < 1000; i++ {
            if check_repeat(get_hash(i+index), c, 5) {
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

