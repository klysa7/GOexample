package main

import (
    /*"errors"*/
    "fmt"
    /*"time"*/
)



func main() {

message:=make(chan string)

go func(){message<-"ping"}()

msg:= <-message
fmt.Println(msg)

}