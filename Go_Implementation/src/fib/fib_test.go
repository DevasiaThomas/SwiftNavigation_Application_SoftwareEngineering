package fib

import "testing"
import "reflect"

//Testprime : Tests the Fib function
func TestFib(t *testing.T) {
	t.Log("Testing if Fib works")
	if Fib(-1) != nil {
		t.Error("Should be nil for zero elements")
	}
	if Fib(0) != nil {
		t.Error("Should be nil for zero elements")
	}
	if !reflect.DeepEqual(Fib(1), []int{0}) {
		t.Error("Should return [0] but returned ", Fib(1))
	}
	if !reflect.DeepEqual(Fib(6), []int{0, 1, 1, 2, 3, 5}) {
		t.Error("Should return [0] but returned ", Fib(6))
	}
}
