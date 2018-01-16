package main

import "fmt"

func decompressed_size(data string) int64 {
    var r int64
    var count, repeat int

    for i := 0; i < len(data); i++ {
        if data[i] == '(' {
            fmt.Sscan(data[i+1:], &count)
            for data[i] != 'x' { i++ }
            fmt.Sscan(data[i+1:], &repeat)
            for data[i] != ')' { i++ }
            r += int64(repeat) * decompressed_size(data[i+1:i+count+1])
            i += count
        } else {
            r++
        }
    }
    return r
}

func main() {
    var data string

    _, err := fmt.Scan(&data)
    if err == nil {
        fmt.Println("size", decompressed_size(data))
    } else {
        fmt.Println("err", err)
    }
}
