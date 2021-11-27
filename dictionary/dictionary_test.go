package dictionary

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and we want %q", got, want)
	}
}
func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err != want {
		t.Fatal("expected to get error")
	}
}
func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("testing for known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("testing for unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		assertError(t, err, nil)
		assertString(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("testing for known word", func(t *testing.T) {
		dictionary := Dictionary{}
		want := "this is just a test"
		dictionary.Add("test", "this is just a test")
		got, err := dictionary.Search("test")
		assertString(t, got, want)
		if err != nil {
			t.Fatal("We shoud not get error", err)
		}
	})
}
