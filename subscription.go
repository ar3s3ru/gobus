package gobus

import "reflect"

func (s *Subscription) GetListeners(typ reflect.Type) (IListenerSet, error) {
    val, ok := (*s)[typ.String()]
    if !ok {
        return nil, ListenersNotFoundErr
    }

    return val, nil
}

func (s *Subscription) AddListener(listener interface{}) (*Subscription) {
    content, err := checkListener(listener)
    if err != nil {
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
