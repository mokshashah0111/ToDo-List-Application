package main

import (
	"gorm.io/gorm"
	"fmt"
)


func validatePassword(password string, confirmPassword string) bool{
	if password == confirmPassword{
		return true
	} else{
		 return false
	}
}
func validCredentials(username, password string) bool {
	var user UserTb
	err := db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false // Credentials are invalid
		}
		// Handle unexpected database errors
		fmt.Println("Database error:", err)
		return false
	}
	return true // Credentials are valid
}
func (UserTb)TableName() string{
	return "UserTb"
}

func (TaskTb)TableName()string{
	return "TaskTb"
}
func checkUserExistence(username string)(bool,error){
	var count int64
	err:= db.Model(&UserTb{}).Where("Username = ?" , username).Count(&count).Error
	if err!=nil{
		return false,err
	}else {
		return count>0,nil
	}
}