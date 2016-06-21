package main

import (
    "fmt"
    "github.com/ar3s3ru/gobus"
)

func printSomething(event string) {
    fmt.Printf("Got this string: %s\n", event)
}

func printSomethingElse(integer int) {
    fmt.Printf("Got this int: %d\n", integer)
}

func main() {
    bus := gobus.NewEventBus(5)
    defer bus.Destruct()

    bus.Subscribe(printSomething, printSomethingElse)
    bus.Publish("Hello world!").    // Publishing a string will call printSomething(string)
    Publish(14).                // Publishing an int will call printSomethingElse(int)
    Publish(10)
}
