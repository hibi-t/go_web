package main

import (
    "fmt"
)

type Event struct {
    Greeter Greeter
}

type Greeter struct {
    Message Message
}

type Message string

type Response struct {
    Writer Writer
}

type Writer string


func NewEvent(g Greeter) (Event) {
    return Event{Greeter: g}
}

func NewGreeter(m Message) Greeter {
    return Greeter{Message: m}
}

func NewMessage(p string) Message {
    return Message(p)
}

func NewResponse(w Writer) (Response) {
    return Response{Writer: w}
}
func NewWriter(r string) Writer {
    return Writer(r)
}

func (e Event) Start() {
    msg := e.Greeter.Great()
    fmt.Println(msg)
}

func (g Greeter) Great() Message {
    return g.Message
}

func (f Response) Res() Writer {
    return f.Writer
}

func main() {
    e := InitializeEvent("takutakutakujiro")
    f := ResponseEvent("test")

    e.Start()
    f.Res()
}