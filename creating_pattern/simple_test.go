package creatingpattern

import "testing"

func TestNewAPI1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("xiaochang")
	if s != "Hi,xiaochang" {
		t.Fatalf("NewAPI1 test fail")
	}
}

func TestNewAPI2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("xiaochang")
	if s != "Hello,xiaochang" {
		t.Fatalf("NewAPI2 test fail")
	}
}
