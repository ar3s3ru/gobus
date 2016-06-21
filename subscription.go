package gobus

import "reflect"

// Returns all the listeners associated to the event type typ, or returns a ListenersNotFoundErr.
func (s *Subscription) GetListeners(typ reflect.Type) (IListenerSet, error) {
    val, ok := (*s)[typ.String()]
    if !ok {
        return nil, ListenersNotFoundErr
    }

    return val, nil
}

// Adds a new listener to the IListenerSet into the Subscription map.
func (s *Subscription) AddListener(listener interface{}) (*Subscription) {
    // Check if the listener is good
    content, err := checkListener(listener)
    if err != nil {
        // If not, panic with ListenerInvalidErr
        panic(err)
    }

    val, ok := (*s)[content.String()]

    if !ok {
        (*s)[content.String()] = &ListenerSet{
            listeners: []interface{}{listener},
        }
    } else {
        val.Add(listener)
    }

    return s
}

// Remove a listener from the IListenerSet into the Subscription map.
func (s *Subscription) RemoveListener(listener interface{}) (*Subscription) {
    content, err := checkListener(listener)
    if err != nil {
        panic(err)
    }

    val, ok := (*s)[content.String()]

    if ok {
        val.Remove(listener)
    }

    return s
}

// Checks if the listener passed is valid.
// If valid, returns a reflect.Type object of the listener input argument;
// if not valid, returns a ListenerInvalidErr and a reflect.Type of the listener passed.
func checkListener(listener interface{}) (reflect.Type, error) {
    lisVal := reflect.TypeOf(listener)

    if lisVal.Kind() != reflect.Func || // Listener must obviously be a function
        lisVal.NumIn()  != 1 || // Listener function must take only 1 input argument
        lisVal.NumOut() != 0 {  // Listener must have no returning argument
        // Listener interface not valid
        return lisVal, ListenerInvalidErr
    }

    // Returns input value kind
    return lisVal.In(0), nil
}
