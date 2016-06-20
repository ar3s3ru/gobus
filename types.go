package gobus

import "sync"

type (
    Subscription map[string]IListenerSet
    Listener     func(interface{})

    EventBus struct {
        dispatcher   chan interface{}
        quit         chan bool
        subscription Subscription
        waitGroup    sync.WaitGroup
    }

    ListenerSet struct {
        listeners []Listener
    }

    IListenerSet interface {
        Add(listener Listener)    IListenerSet
        Remove(listener Listener) IListenerSet
        Values() []Listener
        Empty()  bool
    }
)
