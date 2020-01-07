package multiples_3_5

func SumMultiples(start int, end int) int {
  var sum int

  for i := start; i < end; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }

  return sum
}
