package database

import "gorm.io/gorm"

// DBconn is the global variable that holds the database connection

var (
	DBconn *gorm.DB
)
