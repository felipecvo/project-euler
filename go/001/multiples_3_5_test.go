package multiples_3_5

import (
  "fmt"
  "testing"
)

func TestMultipleSum5To10(t *testing.T) {
  multiples_sum := SumMultiples(1, 10)
  expected := 23

  if multiples_sum != expected {
    t.Errorf("got %d want %d", multiples_sum, expected)
  }
}

func TestEuler(t *testing.T) {
  multiples_sum := SumMultiples(1, 1000)

  fmt.Printf("The result is %d\n", multiples_sum)
}
