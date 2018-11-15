package  main

import (
    "fmt"
    "github.com/mediocregopher/radix.v2/redis"
)


func main() {
    conn, err := redis.Dial("tcp", "localhost:6379")
    defer conn.Close()
    if err != nil {
        fmt.Println("Redis connect error: ", err)
    }

    resp := conn.Cmd("SET","testgo","test20181115")
    if resp.Err != nil {
        fmt.Println("redis error: ", resp.Err)
    }


    resp2 := conn.Cmd("EXPIRE", "testgo", 3600)
    if resp2.Err != nil {
        fmt.Println("redis error: ", resp.Err)
    }

}
