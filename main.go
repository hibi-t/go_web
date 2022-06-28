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

func NewEvent(g Greeter) (Event) {
    return Event{Greeter: g}
}

func NewGreeter(m Message) Greeter {
    return Greeter{Message: m}
}

func NewMessage(p string) Message {
    return Message(p)
}

func (e Event) Start() {
    msg := e.Greeter.Great()
    fmt.Println(msg)
}

func (g Greeter) Great() Message {
    return g.Message
}

type Response struct {
    Writer Writer
}

type Writer struct {
    Pointer Pointer
}

type Pointer string

func NewResponse(w Writer) (Response) {
    return Response{Writer: w}
}
func NewWriter(p Pointer) (Writer) {
    return Writer{Pointer: p}
}
func NewPointer(p string) Pointer {
    return Pointer(p)
}


func (r Response) Res() {
    msg := r.Writer.ResWrite()
    fmt.Println(msg)
}
func (w Writer) ResWrite() Pointer {
    return w.Pointer
}

func main() {
    e := InitializeEvent("takutakutakujiro")
    w := ResponseEvent("test")

    e.Start()
    w.Res()
}