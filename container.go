//+build wireinject

package main

import "github.com/google/wire"

func Container(phrase string) Event {
    wire.Build(NewTest)
    return Event{}
}
