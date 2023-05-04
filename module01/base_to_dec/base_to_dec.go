package main

import (
    "math"
    "strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
    total := 0.0
    letters := map[string]float64{"A":10, "B":11, "C":12, "D":13, "E":14, "F": 15}
    var checkLetter float64
    
    pow := 0.0
    for i := len(value) - 1; i >= 0; i-- {
        checkLetter = letters[string(value[i])]
        if checkLetter != 0 {
            total += math.Pow(float64(base), pow) * checkLetter
        } else {
            num, err := strconv.ParseFloat(string(value[i]), 64)
            if err == nil {
                total += math.Pow(float64(base), pow) * num
            }
        }
        pow++
    }

	return int(total)
}
