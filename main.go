package main

import (
    "fmt"
    "net/http"
    "github.com/mitchellh/mapstructure"
    "time"

)

type Message struct{
    Name string `json:"name"`
    Data interface{} `json:"data"`
}

type Channel struct{
    Id string `json:"id"`
    Name string `json:"name"`
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":4000", nil)
}

func addChannel(data interface{}) (error){
    var channel Channel

    err := mapstructure.Decode(data, &channel)
    if err != nil {
        return err
    }
    channel.Id = "1"
    fmt.Println("added channel")
    return nil
}

func subscribeChannel(socket *websocket.Conn){
    //TODO: rethinkDB query / changefeed
    for {
        time.Sleep(time.Second * 1)
        message := Message{"channel add", 
            Channel{"1", "Software Support"}}
        socket.WriteJSON(message)
        fmt.Println("sent new channel")
    }
}