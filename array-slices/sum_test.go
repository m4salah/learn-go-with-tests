package arrayslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expexted := 15

		if got != expexted {
			t.Errorf("we expected %d but we got %d, %v", expexted, got, numbers)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expexted := 6

		if got != expexted {
			t.Errorf("we expected %d but we got %d, %v", expexted, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("collection of two numbers numbers", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAll(numbers1, numbers2)
		expexted := []int{3, 9}

		if !reflect.DeepEqual(got, expexted) {
			t.Errorf("we expected %d but we got %d, %v, %v", expexted, got, numbers1, numbers2)
		}
	})

	t.Run("collection of two numbers numbers", func(t *testing.T) {
		numbers1 := []int{1, 1, 1}

		got := SumAll(numbers1)
		expexted := []int{3}

		if !reflect.DeepEqual(got, expexted) {
			t.Errorf("we expected %d but we got %d, %v", expexted, got, numbers1)
		}
	})
}

func TestSumTails(t *testing.T) {

	checkSum := func(t testing.TB, got, went []int) {
		t.Helper()
		if !reflect.DeepEqual(got, went) {
			t.Errorf("got %v went %v", got, went)
		}
	}
	t.Run("collection of two numbers numbers", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAllTails(numbers1, numbers2)
		expexted := []int{2, 9}

		checkSum(t, got, expexted)
	})

	t.Run("collection of two numbers numbers", func(t *testing.T) {
		numbers1 := []int{1, 1, 1}

		got := SumAllTails(numbers1)
		expexted := []int{2}

		checkSum(t, got, expexted)

	})

	t.Run("collection of two numbers numbers", func(t *testing.T) {
		numbers1 := []int{}

		got := SumAllTails(numbers1)
		expexted := []int{0}

		checkSum(t, got, expexted)

	})
}
