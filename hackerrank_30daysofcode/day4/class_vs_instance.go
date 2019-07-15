/*
Write a Person class with an instance variable, , and a constructor that takes an integer, , as a parameter. The constructor must assign  to  after confirming the argument passed as  is not negative; if a negative argument is passed as , the constructor should set  to  and print Age is not valid, setting age to 0.. In addition, you must write the following instance methods:

yearPasses() should increase the  instance variable by .
amIOld() should perform the following conditional actions:
If , print You are young..
If  and , print You are a teenager..
Otherwise, print You are old..
To help you learn by example and complete this challenge, much of the code is provided for you, but you'll be writing everything in the future. The code that creates each instance of your Person class is in the main method. Don't worry if you don't understand it all quite yet!
*/

package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	if p.age >= 0 {
		initialAge = p.age
	} else {
		initialAge = 0
		fmt.Println("Age is not valid, setting age to 0.")
	}
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}

}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	p.age++
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
