package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if strings.Compare(repeated, expected) != 0 {
			t.Errorf("expected %q but we got %q", expected, repeated)
		}
	}

	t.Run("repeate a 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeate a n times", func(t *testing.T) {
		repeated := Repeat("a", 9)
		expected := "aaaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}
