package main

import "log"

type User struct{
	FirstName string 
	LastName string
}
 
func main() {
		// Wrong way to create a map
	// var myOtherMap map[string]string
	
								// String
	// myMap := make(map[string]string)

	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie" 

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])


								// int
	// myMap := make(map[string]int)
	// myMap["First"] = 1
	// myMap["Second"] = 2

	// log.Println(myMap["First"])
	// log.Println(myMap["Second"])


								// Types (Structs)
	myMap := make(map[string]User)

	me := User {
		FirstName: "Trevor",
		LastName: "Sawler",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)
	log.Println(myMap["me"].LastName)

	var myNewVar float32
	myNewVar = 11.1
}
