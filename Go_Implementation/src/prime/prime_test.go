package prime

import "testing"

//Testprime : Tests Prime function
func TestPrime(t *testing.T) {
	t.Log("Testing if Prime works")
	if Prime(21) {
		t.Error("21 should not be prime")
	}
	if !Prime(11) {
		t.Error("11 Should be Prime")
	}
	if !Prime(101) {
		t.Error("101 Should be Prime")
	}
	if Prime(0) {
		t.Error("0 should not be prime")
	}
}
