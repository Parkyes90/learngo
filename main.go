package main

import (
	"fmt"
	"github.com/parkyes90/learngo/custom_dict"
)

func main() {
	dictionary := custom_dict.Dictionary{"first": "First word"}
	def, err := dictionary.Search("first")

	err2 := dictionary.Add("test", "test2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
	fmt.Println(dictionary, err2)
}
