package main

import (
  "testing"
  "reflect"
  "strings"
)

func TestSum(t *testing.T) {

  t.Run("collection of 5 numbers", func(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}
    
    got := Sum(numbers)
    want := 15
    
    if got != want {
      t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
  })
  
  t.Run("collection of any size", func(t *testing.T) {
    numbers := []int{1, 2, 3}
    
    got := Sum(numbers)
    want := 6
    
    if got != want {
      t.Errorf("got %d want %d given, %v", got, want, numbers)
    }
  })
  
}

func TestSumAll(t *testing.T) {
  
  got := SumAll([]int{1, 2}, []int{0, 9})
  want := []int{3, 9}
  
  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}

func TestSumAllTails(t *testing.T) {

  checkSums := func(t testing.TB, got, want []int) {
    t.Helper()
    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v want %v", got, want)
    }
  }

  t.Run("make the sums of some slices", func(t *testing.T) {
    got := SumAllTails([]int{1, 2}, []int{0, 9})
    want := []int{2, 9}
    checkSums(t, got, want)
  })
  
  t.Run("safely sum empty slices", func(t *testing.T) {
    got := SumAllTails([]int{}, []int{3, 4, 5})
    want := []int{0, 9}
    checkSums(t, got, want)
  })

}

func TestReduce(t *testing.T) {
  t.Run("multiplication of all elements", func(t *testing.T) {
    multiply := func(x, y int) int {
      return x * y
    }

    AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
  })

  t.Run("concatenate strings", func(t *testing.T) {
    concatenate := func(x, y string) string {
      return x + y
    }

    AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
  })
}

func TestFind(t *testing.T) {
  t.Run("find first even number", func(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    firstEvenNumber, found := Find(numbers, func(x int) bool {
      return x % 2 == 0
    })
    AssertTrue(t, found)
    AssertEqual(t, firstEvenNumber, 2)
  })

  t.Run("find the best programmer", func(t *testing.T) {
    people := []Person{
      Person{Name: "Kent Beck"},
      Person{Name: "Martin Fowler"},
      Person{Name: "Chris James"},
    }

    king, found := Find(people, func(p Person) bool {
      return strings.Contains(p.Name, "Chris")
    })

    AssertTrue(t, found)
    AssertEqual(t, king, Person{Name: "Chris James"})
  })
}

type Person struct {
  Name string
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
