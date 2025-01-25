package main

import (

	"testing"
)
func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected string
	}{
		{
			input: " hello world  ",
		expected: "hello",
		},
	}
	for _,c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if actual!=expected {
			t.Errorf("got %s want %s",actual,expected)
		}

		// for i := range actual {
		// 	word:= actual[i]
		// 	expectedWord := c.expected[i]
		// 	if word != expectedWord {
		// 		t.Errorf("got %s want %s",word,expectedWord)
		// 	}
		// }
	}
}