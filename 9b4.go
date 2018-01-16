package main

import (
    "fmt"
    "strings"
)

func decompressed_size(data string) int64 {
    var r int64
    var count, repeat int
    var tmp string
    var scanning bool

    for i := 0; i < len(data); i++ {
        switch data[i] {
            case '(': {
                tmp = ""
                scanning = true
            }
            case 'x': { // works because of lowercase
                fmt.Sscan(tmp, &count)
                tmp = ""
                scanning = true
            }
            case ')': {
                fmt.Sscan(tmp, &repeat)
                r += int64(repeat) * decompressed_size(data[i+1:i+count+1])
                i += count
                scanning = false
            }
            default: {
                if scanning {
                    tmp += string(data[i])
                } else { r++ }
            }
        }
    }
    return r
}

func main() {
    var field string
    var list []string

    for {
        _, err := fmt.Scan(&field)
        if err != nil { break }
        list = append(list, field)
    }
    fmt.Println("size", decompressed_size(strings.Join(list, "")))
}
