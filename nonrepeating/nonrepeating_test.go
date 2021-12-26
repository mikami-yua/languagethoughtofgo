package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, test := range tests {
		actual := lengthOfNonRepeatingSubStr(test.s)
		if actual != test.ans {
			t.Errorf("got %d for input %s; "+
				"expected %d",
				actual, test.s, test.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ { //13次指数增长
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8
	b.ResetTimer() //重新测试时间

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, s, ans)
		}
	}
}
