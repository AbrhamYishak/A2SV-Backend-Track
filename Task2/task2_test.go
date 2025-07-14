
package main

import "testing"

func mapsEqual(a, b map[string]int) bool {
    if len(a) != len(b) {
        return false
    }
    for k, v := range a {
        if bv, ok := b[k]; !ok || bv != v {
            return false
		
        }
    }
    return true
}
func TestTask2(t *testing.T) {
	palindroemCases := []struct {
		test string
		expect bool
	}{
		{test:"ana",expect: true},
		{test:"test",expect: false},
	}
    counterCases := []struct {
		test string
		expect map[string]int
	}{
		{test:"ana",expect: map[string]int{"a":2,"n":1}},
		{test:"test",expect: map[string]int{"t":2,"s":1,"e":1}},
	}
	for _, c := range palindroemCases {
		got := palindrome(c.test)
		if got != c.expect {
			t.Errorf("palindrome(%v) = %v; want %v", c.test, got, c.expect)
		}
	}
	for _, c := range counterCases {
		got := counter(c.test)
		if !mapsEqual(got, c.expect){
			t.Errorf("counter(%v) = %v; want %v",c.test, got, c.expect)
		}
	}
	
}
