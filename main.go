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

type Test string

func NewEvent(g Greeter) (Event) {
    return Event{Greeter: g}
}

func NewGreeter(m Message) Greeter {
    return Greeter{Message: m}
}

func NewMessage(p string) Message {
    return Message(p)
}

func NewTest(t string) Test {
    return Test(t)
}

func (e Event) Start() {
    msg := e.Greeter.Great()
    fmt.Println(msg)
}

func (g Greeter) Great() Message {
    return g.Message
}

func main() {
    e := InitializeEvent("takutakutakujiro")
    f := Container("test")
    fmt.Println(f)

    e.Start()
}