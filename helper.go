package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"golang.org/x/crypto/bcrypt"

)
func SaveToFile(tasks []TaskTb){
	file,err := os.Create("tasks.txt")

	if err!= nil {
		fmt.Println("Error creating file!")
		return
	}
	defer file.Close()

	for _,task := range tasks{
		status := " "
		if task.Completed{
			status = "x"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status,task.TaskText))
	}
	fmt.Println("Tasks saved to the file successfully")
	askMenu()
}

func showMenu(){
	fmt.Println("\nMenu:")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add task")
	fmt.Println("3. Mark task as completed")
	fmt.Println("4. Save the tasks to file")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}


func showTasks(showPrompt bool){
	// var tasks []TaskTb
	err := db.Where("userID = ?", currentUser.ID).Find(&currentUser.Tasks).Error
	if err!=nil{
		fmt.Println("Error fetching the tasks", err)
		return
	}
	// fmt.Printf("Fetched tasks for userID %d: %+v\n", currentUser.ID, currentUser.Tasks)

	if(len(currentUser.Tasks) == 0){
		fmt.Println("No tasks to show")
	}else {
		fmt.Println("Tasks: ")
		for i, task := range currentUser.Tasks{
			status := " "
			if task.Completed{
				status = "x"
			}
			fmt.Printf("%d. [%s] %s\n",i+1, status,task.TaskText)
		}
	}
	if showPrompt{
		for {
			input := getUserInput("Do you want to add new task? (y/n)")
				switch input{
					case "Y","y":
						AddTasks()
						return
					case "N","n":
						askMenu()
						break
					default:
						fmt.Println("Enter valid answer!")
				
				}
			}
	}
}

func AddTasks(){
	if currentUser.ID == 0 {
        fmt.Println("Error: No user is logged in. Please sign in first.")
        return
    }
	tasktext := getUserInput("Enter the task to be added: ")
	newTask := TaskTb{UserID: currentUser.ID, TaskText: tasktext, Completed: false}
	err:=db.Create(&newTask).Error
	if err != nil {
		fmt.Println("Error adding the task: ", err)
		return
	}
	fmt.Println("Task added !")
	showTasks(true)
}

func parseIndex(input string)(int,error){
	return strconv.Atoi(input)
}

func MarkTaskCompleted(){
	for{
		showTasks(false)
		taskIndex := getUserInput("Enter the number of the task to be marked: ")
		index,err := parseIndex(taskIndex)

		if err!=nil || index<1 || index>len(currentUser.Tasks){
			fmt.Println("Invalid index")
			continue
		}
		task := &currentUser.Tasks[index-1]
		if task.Completed{
			fmt.Println("Task already marked.")
		}else{
			task.Completed= true
			err = db.Save(task).Error
			if err!=nil{
				fmt.Println("Error updating the task: ", err)
				return
			}else {
				fmt.Println("Task marked")
				showTasks(false)
				break
			}
		}
	}
}

func askMenu(){
	for{
		input := getUserInput("Do you want to see the menu?(y/n)")
		switch input{
			case "Y","y":
				// showMenu()
				// option := getUserInput("Enter your choice: ")
				startApp()
			case "N","n":
				showTasks(false)
				os.Exit(0)
			default:
				fmt.Println("Invalid input, please try again")
		}
	}
}

func SignUp() bool{
	var username,password,confirmPassword string
	for{
		username = getUserInput("Enter username: ")
		userExistence,error:= checkUserExistence(username)
		if error!=nil{
			fmt.Println("Error checking the user existence: ",error)
			return false
		}
		if(userExistence){
			fmt.Println("Username already exists! Please try again")
		}else{
			break
		}
	}
	
	for{
		password = getUserInput("Enter password: ")
		confirmPassword = getUserInput("Confirm the password: ")
		if !validatePassword(password,confirmPassword){
			fmt.Println("Passwords do not match. Please try again")
		} else{
			break
		}
	}
	newUser := UserTb{Username: username, Password: password}
	err := db.Create(&newUser).Error
	if err!= nil {
		fmt.Println("Error creating user: ", err)
		return false
	}
	db.Where("username=?",username).First(&currentUser)

	fmt.Println("Account created successfully. Please Sign in to continue")

	return true
}

func hashPassword(password string)(string,error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password) , bcrypt.DefaultCost)
	return string(bytes),err
}

func SignIn() {
    for {
        username := getUserInput("Enter username: ")
        password := getUserInput("Enter password: ")

        var user UserTb
        err := db.Where("username = ? AND password = ?", username, password).First(&user).Error
        if err != nil {
            fmt.Println("Invalid username or password. Please try again!")
            continue
        }
        currentUser = user
        fmt.Println("Sign-in successful! Welcome,", username)
        break
    }
}


func startApp(){
	for {
		showMenu()
		option := getUserInput("Enter your choice: ")
		switch option{
			case "1":
				showTasks(true)
			case "2":
				AddTasks()
			case "3":
				MarkTaskCompleted()
			case "4":
				fmt.Println("Exiting the application.....")
				showTasks(false)
				os.Exit(0)
			default:
				fmt.Println("Invalid input")
		}
		askMenu()
	}
}
 