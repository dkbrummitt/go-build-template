package main

import (
	"errors"
	"fmt"
)

var (
	DebugMode = true                             // public variable, visible outside of package
	seed      = "-02349iyorhkgnmvki009eioryghj8" // not visible outside of package
)

const (
	Pi = 3.1452
)

//show Dog type OMIT
type Dog struct {
	Name  string `json:"name,omitempty"`
	Breed string `json:"breed,omitempty"`
	Age   int    `json:"age,omitempty"`
}

//end show Dog type OMIT

//show Pet type OMIT
type Pet struct {
	Dog //embedded/inherited struct
}

//end show Pet type OMIT

//show hasDog OMIT
func hasDog(d Dog) (yes bool) {
	if d != (Dog{}) {
		yes = true
	}
	return
}

//end show hasDog OMIT

// String method that an read data in dog but cant change anthing.
// and is also thread safe
func (d Dog) String() string {
	if d != (Dog{}) {
		d.Name = "Woot"
		return fmt.Sprintf("My name is %s", d.Name)
	}

	return fmt.Sprintf("So this is embarrasing, I dont exist yet")
}

// ChangeName method that an read/write data in dog but cant change anthing.
// and is also NOT thread safe
func (d *Dog) ChangeName(newName string) (err error) {
	if d != nil {
		d.Name = newName
	}

	err = errors.New("you need to create a dog first")
	return
}

//show main func OMIT
func main() {
	fmt.Println("Proof of life!")
	myDog := Dog{"Rum", "Chorkie", 3}
	myOtherDog := Dog{Name: "Kaiba", Age: 17}

	fmt.Printf("%s %s %d\n", myDog.Name, myDog.Breed, myDog.Age)
	fmt.Printf("All of it g\n", myDog)
	fmt.Printf("Forgetting anyone? %s\n", myOtherDog.Name)

	fmt.Println("myDog says ", myDog.String())
	fmt.Println("is it still Woot?", myDog.Name)
	myOtherDog.ChangeName("Spot")
	fmt.Println("myOtherDog says ", myOtherDog.String())
	fmt.Println("Done!")
}

//end show main func OMIT
