package main

import ( 
	"fmt"
	"strings"
)

var tasks = []TaskTb{}
func main(){
	initDatabase()

	for{
		option := getUserInput("SignIn if  you are an existing user or SignUp if you are a new user: ")
		
		switch strings.ToLower(option){
			case "signin":
				//todo
				SignIn()
				startApp()
				return
			case "signup":
				//todo
				SignUp()
				SignIn()
				startApp()
				return
			default:
				fmt.Println("Invalid option, try again")

		}

	}
}










