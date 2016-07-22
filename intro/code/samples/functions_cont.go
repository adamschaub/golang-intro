// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("*Knock knock* Who's there?")

	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}

	if thatsMe := AmITheOneWhoKnocks(name); thatsMe == true {
		fmt.Println("I am the danger.")
	} else {
		fmt.Println("Who?")
	}
}

func AmITheOneWhoKnocks(name string) bool {
	return name == "Walter"
}
