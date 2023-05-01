package main

import "testing"

func TestNumInList(t *testing.T) {
    var inList bool
    var expected bool

    inList = NumInList([]int{1,2,3,4,5}, 5) // true
    expected = true
    if inList != expected {
        t.Errorf("Expected '%t' but got '%t'", expected, inList)
    }

    inList = NumInList([]int{3}, 5) // false
    expected = false
    if inList != expected {
        t.Errorf("Expected '%t' but got '%t'", expected, inList)
    }

    inList = NumInList([]int{3,5,3,5,3}, 5) // true
    expected = true
    if inList != expected {
        t.Errorf("Expected '%t' but got '%t'", expected, inList)
    }

    inList = NumInList([]int{4,2,22,-10,8}, -10) // true
    expected = true
    if inList != expected {
        t.Errorf("Expected '%t' but got '%t'", expected, inList)
    }

    inList = NumInList(nil, 5) // false
    expected = false
    if inList != expected {
        t.Errorf("Expected '%t' but got '%t'", expected, inList)
    }

    inList = NumInList([]int{}, 5) // false
    expected = false
    if inList != expected {
        t.Errorf("Expected '%t' but got '%t'", expected, inList)
    }
}
