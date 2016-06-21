package gobus

import "reflect"

// Add a new listener into the ListenerSet only if it's actually a new listener.
// Returns a ListenerSet pointer for method chaining pattern.
func (set *ListenerSet) Add(listener interface{}) IListenerSet {
    if _, ok := set.hasListener(listener); !ok {
        set.listeners = append(set.listeners, listener)
    }

    return set
}

// Remove a listener from the ListenerSet using a new slice of listeners.
// Returns a ListenerSet pointer for method chaining pattern.
//
// N.B. the append() method could be really performance bad...
func (set *ListenerSet) Remove(listener interface{}) IListenerSet {
    if i, ok := set.hasListener(listener); ok {
        set.listeners = append(set.listeners[:i], set.listeners[i+1:]...)
    }

    return set
}

// Returns all the values of the ListenerSet.
func (set *ListenerSet) Values() []interface{} {
    return set.listeners
}

// Checks if the ListenerSet is empty.
func (set *ListenerSet) Empty() bool {
    return len(set.listeners) == 0
}

// Performs a shallow equality on the pointer value of the listener against all the listeners into the ListenerSet.
// If a match is found, returns (index-of-the-match, true);
// if not, returns (-1, false).
func (set *ListenerSet) hasListener(listener interface{}) (int, bool) {
    findPointer := reflect.ValueOf(listener).Pointer()

    for i, mListener := range set.listeners {
        if reflect.ValueOf(mListener).Pointer() == findPointer {
            return i, true
        }
    }

    return -1, false
}