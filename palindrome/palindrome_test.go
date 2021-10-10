package main

import "testing"

func TestPalindrome(t *testing.T) {
	t.Run("Should be pass a valid word", func(t *testing.T) {
		const word string = "tenet"
		isValid := Palindrome(word)

		if !isValid {
			t.Errorf("\nerror, the word is not a palindrome")
		}
	})

	t.Run("Should be pass a invalid word", func(t *testing.T) {
		const word string = "teste"
		isValid := Palindrome(word)

		if isValid {
			t.Errorf("\nError, the word is a palindrome")
		}
	})
}
