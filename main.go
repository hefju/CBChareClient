package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc/jsonrpc"
)

type Args struct {
    Who string
}
type Reply struct {
    Message string
}

func main() {
    //fmt.Println("from server:")

    client, err := net.Dial("tcp", "127.0.0.1:8081")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    args := &Args{"caonima"}
    reply := new(Reply)
    c := jsonrpc.NewClient(client)
    err = c.Call("Hello.Say", args, reply)
    if err != nil {
        log.Fatal("mygod:", err)
    }
    fmt.Println("from server:", reply)
}
