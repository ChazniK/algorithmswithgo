package main

import "testing"

func TestSum(t *testing.T) {
    var sum int
    var expected int

    sum = Sum([]int{1,2,3,4,5}) // 15
    expected = 15
    if sum != expected {
        t.Errorf("Expected '%d' but got '%d'", expected, sum)
    }

    sum = Sum([]int{3,5,3,5,3}) // 19
    expected = 19
    if sum != expected {
        t.Errorf("Expected '%d' but got '%d'", expected, sum)
    }

    sum = Sum([]int{}) // 0
    expected = 0
    if sum != expected {
        t.Errorf("Expected '%d' but got '%d'", expected, sum)
    }

    expected = 0
    sum = Sum(nil) // 0
    if sum != expected {
        t.Errorf("Expected '%d' but got '%d'", expected, sum)
    }
}

