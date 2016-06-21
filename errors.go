package gobus

import "errors"

var (
    ListenersNotFoundErr = errors.New("Cannot found listeners, check for unhandled event subscriptions")
    ListenerInvalidErr   = errors.New("Invalid listener passed through, must be unary function with no return argument")
)
