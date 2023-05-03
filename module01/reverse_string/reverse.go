package main

//Given a string, return its reverse.
//Ex.
//Reverse("cat") // "tac"

func Reverse(word string) string {

    reversed := ""
    for i := len(word) - 1; i >= 0; i-- {
        reversed += string(word[i])
    }

    return reversed
}
