package main

import (
	"fmt"

	"github.com/Tomoya185-miyawaki/travel-buddy/db"
	"github.com/Tomoya185-miyawaki/travel-buddy/domain"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Done migrate")
	defer db.CloseDB(dbConn)
	db.NewDB().AutoMigrate(&domain.User{})
}
