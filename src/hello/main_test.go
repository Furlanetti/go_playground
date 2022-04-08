package main

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello World!"
    if "Hello World!" != expected {
        t.Errorf("should")
    }
}
