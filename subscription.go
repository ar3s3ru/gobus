package gobus

import (
    "reflect"
)

func (s *Subscription) GetListeners(kind reflect.Kind) (IListenerSet, error) {
    val, ok := (*s)[kind.String()]
    if !ok {
        //return nil, ERRORNOTFOUND
    }

    return val, nil
}

func (s *Subscription) AddListener(kind reflect.Kind, listener Listener) (*Subscription) {
    val, ok := (*s)[kind.String()]

    if !ok {
        (*s)[kind.String()] = &ListenerSet{
            listeners: []Listener{listener},
        }
    } else {
        val.Add(listener)
    }

    return s
}

func (s *Subscription) RemoveListener(kind reflect.Kind, listener Listener) (*Subscription) {
    val, ok := (*s)[kind.String()]

    if ok {
        val.Remove(listener)
    }

    return s
}