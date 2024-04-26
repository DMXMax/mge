package main

import (
	"example/mge/util"
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//create a function called main that generates a random result from the Actions map and a random result from the subject map and prints them out
//hint: use the rand package
//hint: use the len function
//hint: use the rand.Intn function
//hint: use the fmt.Println function
//hint: use the util.Action map
//hint: use the util.Subject map

func main() {
	//get a random Action

	event := util.GetEvent(10)
	fmt.Println(cases.Title(language.AmericanEnglish).String(event.String()))

	//get a random Subject
	//print out the results

}

/*func properTitle(input string) string {
	words := strings.Split(input, " ")
	smallwords := " a an on the to "

	for index, word := range words {
		if strings.Contains(smallwords, " "+word+" ") && word != string(word[0]) {
			words[index] = word
		} else {
			words[index] = cases.Title(word)
		}
	}
	return strings.Join(words, " ")
}*/
