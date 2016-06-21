package gobus

import "reflect"

func (set *ListenerSet) Add(listener interface{}) IListenerSet {
    if _, ok := set.hasListener(listener); !ok {
        set.listeners = append(set.listeners, listener)
    }

    return set
}

func (set *ListenerSet) Remove(listener interface{}) IListenerSet {
    if i, ok := set.hasListener(listener); ok {
        set.listeners = append(set.listeners[:i], set.listeners[i+1:]...)
    }

    return set
}

func (set *ListenerSet) Values() []interface{} {
    return set.listeners
}

func (set *ListenerSet) Empty() bool {
    return len(set.listeners) == 0
}

func (set *ListenerSet) hasListener(listener interface{}) (int, bool) {
    findPointer := reflect.ValueOf(listener).Pointer()

    for i, mListener := range set.listeners {
        if reflect.ValueOf(mListener).Pointer() == findPointer {
            return i, true
        }
    }

    return -1, false
}