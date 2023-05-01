package main

// Given a list of numbers, determine if a specific number is in that list.
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

    for _, currentNum := range list {
        if currentNum == num {
            return true
        } else {
            continue
        }
    }

    return false
}
