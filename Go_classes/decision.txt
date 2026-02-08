package main

import (
	"log"
)

func main() {

	myVar := "ree"

	switch myVar {

	case "cat":
		log.Println("Cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")


	case "fish":
		log.Println("cat is set to fish")
		
	default:
		log.Println("Cat is Something else")
	}
	
	// If Statements

	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue{
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue{
		log.Println("1")
	}else{
		log.Println("an else statements")
	}

	cat := "cat"

	if cat == "cat"{
		log.Println("Cat is cat")
	}else{
		log.Println("Cat is not cat")
	}

	if state 1
	var isTrue bool

	isTrue = true

	if isTrue == true{
		log.Println("isTrue is", isTrue)
	}else{
		log.Println("isTrue is", isTrue)
	}

}
