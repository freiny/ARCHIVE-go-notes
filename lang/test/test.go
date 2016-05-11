package libtest

import "testing"

func Reverse(s string) string {
	ret := ""
	for i := len(s) - 1; i >= 0; i-- {
		ret += string(s[i])
	}
	return ret
}

var reverseTests = []struct {
	s   string
	out string
}{
	{"", ""},
	{"a", "a"},
	{"abcde", "edcba"},
}

func TestReverse(t *testing.T) {
	for _, test := range reverseTests {
		actual := Reverse(test.s)
		if actual != test.out {
			t.Errorf("Reverse(%q) = %q ; %q wanted", test.s, actual, test.out)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcdefghijklmnopqrstuvwxyz")
	}
}
