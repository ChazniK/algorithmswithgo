//Given a list of numbers, sum them all up and return the sum.

package main

func Sum(numbers []int) int {
    sum := 0

    for _, num := range numbers {
        sum += num
    }
    return sum
}
