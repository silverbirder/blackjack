package sample01

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := HelloWorld("hoge")
	expected := "hello world, hoge"
	if got != expected {
		t.Errorf("got %v\nwant %v", got, expected)
	}
}
