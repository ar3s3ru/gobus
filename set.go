package gobus

import "reflect"

func (set *ListenerSet) Add(listener Listener) IListenerSet {
    if _, ok := set.hasListener(listener); !ok {
        set.listeners = append(set.listeners, listener)
    }

    return set
}

func (set *ListenerSet) Remove(listener Listener) IListenerSet {
    if i, ok := set.hasListener(listener); ok {
        set.listeners = append(set.listeners[:i], set.listeners[i+1:]...)
    }

    return set
}

func (set *ListenerSet) Values() []Listener {
    return set.listeners
}

func (set *ListenerSet) Empty() bool {
    return len(set.listeners) == 0
}

func (set *ListenerSet) hasListener(listener Listener) (int, bool) {
    findPointer := reflect.ValueOf(listener).Pointer()

    for i, mListener := range set.listeners {
        if reflect.ValueOf(mListener).Pointer() == findPointer {
            return i, true
        }
    }

    return -1, false
}