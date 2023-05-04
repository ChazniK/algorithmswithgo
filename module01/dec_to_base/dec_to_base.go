package main

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//

func DecToBase(dec, base int) string {

    letters := map[int]string{10:"A", 11:"B", 12:"C", 13:"D", 14:"E", 15:"F"}
    convertedNum := ""

    for dec > 0 {
        rem := dec % base
        if rem > 9 { 
            convertedNum = letters[rem] + convertedNum
        } else {
            convertedNum = fmt.Sprintf("%d", rem) + convertedNum
        }
        dec /= base
    }

	return convertedNum
}

