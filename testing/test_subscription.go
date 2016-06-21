package testing

import (
    "testing"
    "github.com/ar3s3ru/gobus"
    "reflect"
)

func functionCalling(listener interface{}, event interface{}) {
    funct, evt := reflect.ValueOf(listener), reflect.ValueOf(event)
    funct.Call([]reflect.Value{evt})
}

func TestGet(t *testing.T) {
    sub := make(gobus.Subscription)
    defer func() {
        if r := recover(); r != nil {
            t.Fatal("Panicked with", r)
        }
    }()

    // Using a boolean guard to see if execution is done or not
    executed := false
    function1 := func(str string) {
        executed = true
    }

    // Add function into the Subscription Map, to the reflect.String type
    sub.AddListener(function1)
    // Retrieving string listeners
    list, err := sub.GetListeners(reflect.TypeOf(""))
    if err != nil {
        // No listeners found...
        t.Fatal("Failed with", err)
    } else if list.Empty() || len(list.Values()) != 1 {
        // ...or too much listeners found
        t.Fatal("Wrong size of listener slice")
    }

    // Calling the first (and only) listener here
    functionCalling(list.Values()[0], "")

    if !executed {
        // Listener has not been executed, error
        t.Fatal("function1 not executed, error")
    }

    // Checking empty listeners
    list, err = sub.GetListeners(reflect.TypeOf(1))
    if err == nil || !list.Empty() {
        t.Fatal("Should have error or empty list")
    }
}

//func TestAdd(t *testing.T) {
//    sub := make(gobus.Subscription)
//
//    function1 :=
//}

//func TestGRemove(t *testing.T) {
//    sub := make(gobus.Subscription)
//}
