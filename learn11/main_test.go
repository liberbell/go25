package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid_data", 100.0, 10.0, 10.0, false},
	{"invalid_data", 100.0, 0, 0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but gone one", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)
// 	if err == nil {
// 		t.Error("Dig not get an error when we should not have")
// 	}
// }
