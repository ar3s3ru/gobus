# GoBus [![Build Status](https://travis-ci.org/ar3s3ru/gobus.svg?branch=master)](https://travis-ci.org/ar3s3ru/gobus)

Simple asynchronous, <b>content-based</b> event bus for Go.

## Usage

GoBus provides a straightforward implementation for an Event Bus.</br>
Start using the Event Bus this way:
```go
bus := gobus.NewEventBus()                  // Un-buffered channel
bus := gobus.NewEventBusBuffered(chanSize)  // Buffered channel
defer bus.Destruct()
```
GoBus can use a buffered and an un-buffered channel to dispatch events as they arrive.

Always remember to call  ```bus.Destruct()``` at the end of the Event Bus usage, as it's needed for
cleanup purposes (closing channels, returning asynchronous goroutines, ...).

#### (Un)Subscription

You can subscribe (and unsubscribe) one or more listeners to the Event Bus like this:

```go
func useString(event string) {
    // Do something
}

bus.Subscribe(useString)
bus.UnSubscribe(useString)

// Method chaining
bus.Subscribe(function1).Subscribe(function2).Subscribe(function3).
    UnSubscribe(function2).UnSubscribe(function3)

// Variadic arguments
bus.Subscribe(function1, function2, function3)
bus.UnSubscribe(function1, function3, function2)

// Having fun :-)
bus.Subscribe(function1, function2, function3).
    UnSubscribe(function1, function2, function3)
```

Listeners must be <b>unary procedures</b>, functions with <b>one input argument</b> and <b>no return arguments.</b>

Listeners are grouped together by their input argument types (meaning that publishing a string <b><i>will call every string
listeners registered to the bus</i></b>).

#### Publishing

You can publish events to the Event Bus this way:

```go
bus.Publish("HelloWorld!")

// Method chaining
bus.Publish("HelloWorld!").Publish(12)
```

Events are pushed to a dispatcher channel which will asynchronously calls all the listeners registered
to the event type.

<b>Being asynchronous through goroutines, there are no guarantees on the listeners calling order.</b>

## Contributing

Biggest contribution towards this library is to use it and give us feedback for further improvements and additions.

For direct contributions, branch of from master and do _pull request_.

## License

This library is distributed under the MIT License found in the
[LICENSE](https://github.com/ar3s3ru/gobus/blob/master/LICENSE) file.

Written by Danilo Cianfrone.
