package gobus

import "sync"

type (
    // Map that holds all listener references, indexed through input argument name.</br>
    // Uses an IListenerSet interface as return type, for optimization purposes.</br>
    // Example of subscriptions map:
    //     map
    //     |--> string (built-in)
    //     |    |--> printString1(str string), printString2(str string)
    //     |
    //     |--> Struct1 (user-defined)
    //          |--> printStruct1(s1 Struct1), doSomethingStruct1(s1 Struct1)
    //
    Subscription map[string]IListenerSet

    // EventBus, of course
    EventBus struct {
        dispatcher   chan interface{}
        subscription Subscription
        waitGroup    sync.WaitGroup
    }

    // ListenerSet is a struct that uses an interface{} slice
    // to implement a listeners set (IListenerSet interface).
    ListenerSet struct {
        listeners []interface{}
    }

    // Interface for listener set.
    IListenerSet interface {
        Add(listener interface{})    IListenerSet
        Remove(listener interface{}) IListenerSet
        Values() []interface{}
        Empty()  bool
    }
)
