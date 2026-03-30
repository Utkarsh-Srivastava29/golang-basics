package main

import "fmt"
//we cannot use the walrus operator outside the functions
//like jwtToken:=300000 will throw an error if we declare it outside the main function and that is pretty much oabout the above
//but the following is allowed *var jwtToken=300000*



//variables can be just changed during the runtime but in case of constants they cannot be changed at all 
const LoginToknen string = "sdfjsdfjksdfjsdf"//public constant
func main() {
	var username string = "utkarsh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var small uint8 = 255
	fmt.Println(small)
	fmt.Printf("Variable is of type: %T \n", small)

	var smallFloat float32 = 255.455441234578
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	//we can have precisn for the case of float64 and it will be very much precison and that is it
	//default values and the aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type and implicit way for declaring the variables
	var website = "learncodingonline.in"
	fmt.Println(website)

	//no var style
	//:= walrus operator
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)
     

	fmt.Println(LoginToknen)
	fmt.Println("Variable is of type: ", LoginToknen)
}
