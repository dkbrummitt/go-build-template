package main

import (
	"fmt"
)

//show Dog type OMIT
type Dog struct {
	Name  string `json:"name,omitempty"`
	Breed string `json:"breed,omitempty"`
	Age   int    `json:"age,omitempty"`
}

//end show Dog type OMIT

//show main func OMIT
func main() {
	fmt.Println("Slides demo!")
	//initializaing a struct this way requires all inputs
	myDog := Dog{"Rum", "Chorkie", 3} // HLfull
	//initializaing a struct this allows you to specify what set
	myOtherDog := Dog{Name: "Kaiba", Age: 17} // HLparial
	fmt.Printf("%s %s %d\n", myDog.Name, myDog.Breed, myDog.Age)
	fmt.Printf("All of it %+v\n", myDog)
	fmt.Printf("Forgetting anyone? %s\n", myOtherDog.Name)
	fmt.Println("Done!")
}

//end show main func OMIT
