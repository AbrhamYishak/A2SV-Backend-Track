package main

import "testing"

func TestAverage(t *testing.T) {
	cases := []struct {
		sum    float64
		count  int
		expect float64
	}{
		{100, 4, 25},
		{50, 2, 25},
		{0, 0, 0},
		{7, 3, 7.0/3.0},
	}

	for _, c := range cases {
		got := Average(c.sum, c.count)
		if got != c.expect {
			t.Errorf("average(%v, %v) = %v; want %v", c.sum, c.count, got, c.expect)
		}
	}
	
}
