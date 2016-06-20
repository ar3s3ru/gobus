package gobus

import (
    "reflect"
    "log"
)

func NewEventBus(chanSize int) (*EventBus) {
    bus := &EventBus{
        subscription: make(Subscription),
        dispatcher:   make(chan interface{}, chanSize),
        quit:         make(chan bool),
    }

    go bus.pollerBus()
    return bus
}

func (bus *EventBus) Destruct() {
    bus.waitGroup.Wait()
    bus.quit <- true
}

func (bus *EventBus) Subscribe(kind reflect.Kind, listener Listener) (*EventBus) {
    bus.subscription.AddListener(kind, listener)
    return bus
}

func (bus *EventBus) UnSubscribe(kind reflect.Kind, listener Listener) (*EventBus) {
    bus.subscription.RemoveListener(kind, listener)
    return bus
}

func (bus *EventBus) Publish(event interface{}) (*EventBus) {
    // TODO: event passed by value
    log.Printf("Publishing event %v\n", event)
    bus.waitGroup.Add(1)
    bus.dispatcher <- event
    return bus
}

func (bus *EventBus) alertListeners(event interface{}) {
    listeners, err := bus.subscription.GetListeners(reflect.TypeOf(event).Kind())
    if err == nil {
        for _, listener := range listeners.Values() {
            // DEBUG
            log.Printf("Alerting listener %v\n", listener)

            bus.waitGroup.Add(1)
            go bus.executingWithWaiting(listener, event)
        }
    }

    bus.waitGroup.Done()
}

func (bus *EventBus) executingWithWaiting(listener Listener, event interface{}) {
    // DEBUG
    log.Printf("Executing listener %v with %v\n", listener, event)

    listener(event)
    bus.waitGroup.Done()
}

func (bus *EventBus) pollerBus() {
    for {
        select {
        case v := <-bus.dispatcher:
            // DEBUG
            log.Printf("Received event %v\n", v)
            go bus.alertListeners(v)

        case <-bus.quit:
            // Exiting from looping
            log.Printf("Received quitting\n")
            return
        }
    }
}
