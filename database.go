package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

type TaskTb struct{
	TaskID uint	`gorm:"column:ID;primarykey"`
	UserID uint	`gorm:"column:userID;default:0;not null"`
	TaskText string `gorm:"column:taskText"`
	Completed bool `gorm:"column:Completed"`
	User UserTb	`gorm:"foreignKey:UserID"`
}

type UserTb struct{
	ID uint	`gorm:"primaryKey; column:ID"`
	Username string	`gorm:"unique;not null;column:username"`
	Password string `gorm:"not null;column:password"`
	Tasks []TaskTb `gorm:"foreignKey:UserID"`
}

var db *gorm.DB
var currentUser UserTb

func initDatabase(){
	var err error
	db,err = gorm.Open(sqlite.Open(`C:\Users\smoks\Documents\Interview Topics\ToDoApplication\databases\toDo.db?_busy_timeout=5000`), &gorm.Config{})
	if err!= nil{
		panic("failed to connect database")
	}else{
		sqlDB,_ := db.DB()
		_,err = sqlDB.Exec("PRAGMA foreign_keys = ON;")
		if err!=nil{	
			panic("failed to enable the foreign keys")
		}
		err =db.AutoMigrate(&TaskTb{},&UserTb{})
		if err != nil {
			panic("failed to migrate database schema")
		}
		// fmt.Println("Database connected and migrated succesfully.")
		// defer sqlDB.Close()
	}
}




