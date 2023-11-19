package main

import (
	"fmt"

	mydict "github.com/bys096/dictinary/dict"
)

func main() {
	/*
		Search Dictionary
		dictionary := mydict.Dictionary{"first": "First word"}
		definition, err := dictionary.Search("first")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/

	/*
		Add & Search & Error Handling Dictionary
		dictionary := mydict.Dictionary{}
		word := "hello"
		definition := "Greeting"
		err := dictionary.Add(word, definition)
		if err != nil {
			fmt.Println(err)
		}
		resultWord, _ := dictionary.Search(word)
		fmt.Println("found word:", word, ",definition:", resultWord)

		err2 := dictionary.Add(word, definition)
		if err2 != nil {
			fmt.Println(err2)
		}
	*/

	/*
		dictionary := mydict.Dictionary{}
		baseWord := "hello"
		dictionary.Add(baseWord, "First")
		err := dictionary.Update(baseWord, "Second")
		if err != nil {
			fmt.Println(err)
		}
		word, _ := dictionary.Search(baseWord)
		fmt.Println(word)
	*/
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
