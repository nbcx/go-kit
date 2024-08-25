package k

import (
	"fmt"
	"testing"
	"time"
)

func TestSpeed(t *testing.T) {
	tests := []struct {
		name string
		fs   []func()
		want time.Duration
	}{
		{"test", []func(){func() { time.Sleep(time.Second) }}, time.Second},
		{"test", []func(){func() { time.Sleep(time.Second) }, func() { time.Sleep(2 * time.Second) }}, 3 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Speed(tt.fs...); got < tt.want {
				t.Errorf("Speed() = %v, want %v", got, tt.want)
			} else {
				fmt.Println("got", got, tt.want, got < tt.want)
			}
		})
	}
}
