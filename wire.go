//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) Event {
    wire.Build(NewEvent, NewGreeter, NewMessage)
    return Event{}
}

func ResponseEvent(phrase string) Response {
    wire.Build(NewResponse, NewWriter, NewPointer)
    return Response{}
}