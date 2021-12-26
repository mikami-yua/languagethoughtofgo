package main

import "testing"

//表格驱动测试
func TestCalcTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 0},
		{5, 12, 13},
		{8, 15, 7},
	}

	for _, test := range tests {
		if actual := calcTriangle(test.a, test.b); actual != test.c {
			t.Errorf("calcTriangle(%d %d); "+
				"got %d;expected %d",
				test.a, test.b, actual, test.c)
		}
	}
}
