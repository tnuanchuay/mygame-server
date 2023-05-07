package main

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}
	expect := []int{0, 2, 3, 4, 5}
	v := append(a[:1], a[1+1:]...)
	if !reflect.DeepEqual(expect, v) {
		t.Fail()
	}
}

func TestSlice2(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}
	expect := []int{1, 2, 3, 4, 5}
	v := append(a[:0], a[0+1:]...)
	if !reflect.DeepEqual(expect, v) {
		t.Fail()
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan string, 100)
	chA := getChA(ch)
	chB := getChB(ch)

	if chA != ch {
		t.Fail()
	}

	if chB != ch {
		t.Fail()
	}
}

func getChA(ch chan string) <-chan string {
	return ch
}

func getChB(ch chan string) chan<- string {
	return ch
}
