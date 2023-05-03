package main

import "testing"

func TestReverse(t *testing.T) {
    var reversedString string
    var expected string

    reversedString = Reverse("")
    expected = ""
    if reversedString != expected {
        t.Errorf("Expected '%s' but got '%s'", expected, reversedString)
    }

    reversedString = Reverse("cat") // "tac"
    expected = "tac"
    if reversedString != expected {
        t.Errorf("Expected '%s' but got '%s'", expected, reversedString)
    }

    reversedString = Reverse("ca t") // "t ac"
    expected = "t ac"
    if reversedString != expected {
        t.Errorf("Expected '%s' but got '%s'", expected, reversedString)
    }

    reversedString = Reverse("alphabet") // "tebahpla"
    expected =  "tebahpla"
    if reversedString != expected {
        t.Errorf("Expected '%s' but got '%s'", expected, reversedString)
    }
}
