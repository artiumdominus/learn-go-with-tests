package main

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
  sum := 0
  for _, number := range numbers {
    sum += number
  }
  return sum
}

func SumAll(numbersToSum ...[]int) []int {
  var sums []int
  for _, numbers := range numbersToSum {
    sums = append(sums, Sum(numbers))
  }
  
  return sums
}

// SumAllTails calculates the sums of all but the first number giver a collection of slices.
func SumAllTails(numbers ...[]int) []int {
  sumTail := func(acc, x []int) []int {
    if len(x) == 0 {
      return append(acc, 0)
    } else {
      tail := x[1:]
      return append(acc, Sum(tail))
    }
  }

  return Reduce(numbers, sumTail, []int{})
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
  var result = initialValue
  for _, x := range collection {
    result = f(result, x)
  }
  return result
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
  for _, v := range items {
    if predicate(v) {
      return v, true
    }
  }
  return
}