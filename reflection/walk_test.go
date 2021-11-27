package reflection

import (
	"reflect"
	"testing"
)

type person struct {
	Name    string
	Profile profile
}

type profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		}, {
			"Struct with two string field",
			struct {
				Name string
				City string
			}{"Chris", "Cairo"},
			[]string{"Chris", "Cairo"},
		}, {
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		}, {
			"Nested Field",
			person{
				"Chris",
				profile{33, "Cairo"},
			},
			[]string{"Chris", "Cairo"},
		}, {
			"Pointers thing",
			&person{
				"Chris",
				profile{33, "Cairo"},
			},
			[]string{"Chris", "Cairo"},
		}, {
			"Slices",
			[]profile{
				{33, "Cairo"},
				{34, "Alex"},
			},
			[]string{"Cairo", "Alex"},
		}, {
			"Arrays",
			[2]profile{
				{33, "Cairo"},
				{34, "Alex"},
			},
			[]string{"Cairo", "Alex"},
		}, {
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan profile)

		go func() {
			aChannel <- profile{33, "Cairo"}
			aChannel <- profile{34, "Alex"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Cairo", "Alex"}

		walk(aChannel, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (profile, profile) {
			return profile{33, "Cairo"}, profile{34, "Alex"}
		}

		var got []string
		want := []string{"Cairo", "Alex"}

		walk(aFunction, func(s string) {
			got = append(got, s)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
