//+build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) Event {
    wire.Build(NewEvent, NewGreeter, NewMessage)
    return Event{}
}