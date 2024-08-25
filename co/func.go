package co

import (
	"errors"
	"time"
)

var ErrTimeout = errors.New("go timeout func")

func Func(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f()
	}()
	return ch
}

func Timeout(timeout time.Duration, f func() error) chan error {
	ch := make(chan error)
	go func() {
		var err error
		select {
		case err = <-Func(f):
			ch <- err
		case <-time.After(timeout):
			ch <- ErrTimeout
		}
	}()
	return ch
}

func TimeoutWait(f func() error, timeout time.Duration) (err error) {
	done := make(chan bool)
	go func() {
		err = f()
		done <- true
	}()
	select {
	case <-time.After(timeout):
		return ErrTimeout
	case <-done:
		return
	}
}
