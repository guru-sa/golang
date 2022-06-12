package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Surname string
	Hobbies []string
	id      string
}

type RailroadWideChecker interface {
	CheckRailsWidth() int
}

type Railroad struct {
	Width int
}

type Train struct {
	TrainWidth int
}

func main() {
	println("Hello World!")

	//[Variables and constants]
	//Explicitly declaring a "string" variable
	var explicit string = "Hello, I'm a explicitly declared variable"

	//Implicitly declaring a "string". Type inferred
	inferred := ", I'm an inferred variable "

	fmt.Println("Variable 'explicit' is of type:",
		reflect.TypeOf(explicit))
	fmt.Println("Variable 'inferred' is of type:",
		reflect.TypeOf(inferred))

	//[Flow control]
	ten := 10
	if ten == 20 {
		println("This shouldn't be printed as 10 isn't equal to 20")
	} else {
		println("Ten is not equals to 20")
	}

	if "a" == "b" || 10 == 10 || true == false {
		println("10 is equal to 10")
	} else if 11 == 11 && "go" == "go" {
		println("This isn't print because previous condition was satisfied")
	} else {
		println("In case no condition is satisfied, print this")
	}

	number := 3
	switch number {
	case 1:
		println("Number is 1")
	case 2:
		println("Number is 2")
	case 3:
		println("Number is 3")
	}

	for i := 0; i <= 10; i++ {
		println(i)
	}

	my_array := [5]int{1, 2, 3, 4, 5}
	for index, value := range my_array {
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}

	for index := 0; index < len(my_array); index++ {
		value := my_array[index]
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}

	//[Functions]
	hello("gurusa")

	//anonymous function
	add := func(m int) int {
		return m + 1
	}

	result := add(6)
	println(result)

	//Closures
	addN := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}
	addFive := addN(5)
	result = addFive(6)
	println(result)

	//Function with undermined number of parameters
	fmt.Printf("%d\n", sum(1, 2, 3))
	fmt.Printf("%d\n", sum(4, 5, 6, 7, 8))

	//Creating errors, handling errors and returning errors
	err := doesReturnError()
	if err != nil {
		// panic(err)
	}

	//Slices
	mySlices := make([]int, 10)
	mySlices = append(mySlices, 5)
	mySlices = mySlices[1:]
	mySlices = append(mySlices[:1], mySlices[2:]...)

	//Map
	myJsonMap := make(map[string]interface{})
	jsonData := []byte(`{"hello":"world"}`)
	err = json.Unmarshal(jsonData, &myJsonMap)
	if err != nil {
		// panic(err)
	}
	fmt.Printf("%s\n", myJsonMap["hello"])

	//Pointer
	pointer_to_number := &number
	println(pointer_to_number)
	println(*pointer_to_number)

	p := &Person{
		Name:    "Mario",
		Surname: "Castro",
		Hobbies: []string{"cycling", "electronics", "planes"},
		id:      "sa3-223-asd",
	}
	fmt.Printf("%s likes %s, %s and %s\n", p.GetFullName(), p.Hobbies[0],
		p.Hobbies[1], p.Hobbies[2])

	//interfaces
	railroad := &Railroad{
		Width: 10,
	}
	passengerTrain := &Train{
		TrainWidth: 10,
	}
	cargoTrain := &Train{
		TrainWidth: 15,
	}

	canPassengerTrainPass := railroad.IsCorrectSizeTrain(passengerTrain)
	canCargoTrainPass := railroad.IsCorrectSizeTrain(cargoTrain)

	fmt.Printf("Can passenger train pass? %v\n", canPassengerTrainPass)
	fmt.Printf("Can cargo train pass? %v\n", canCargoTrainPass)

}

func (r *Railroad) IsCorrectSizeTrain(c RailroadWideChecker) bool {
	return c.CheckRailsWidth() == r.Width
}

func (t *Train) CheckRailsWidth() int {
	return t.TrainWidth
}

//Method
func (person *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", person.Name, person.Surname)
}

func hello(message string) error {
	fmt.Printf("Hello %s\n", message)
	return nil
}

func doesReturnError() error {
	err := errors.New("this function simply returns an error")
	return err
}

func sum(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return
}
