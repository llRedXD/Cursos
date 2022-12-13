package main

import "testing"

func TestAdd(t *testing.T){

    got := Add(4, 7)
    want := 10

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}