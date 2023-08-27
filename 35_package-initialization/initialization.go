package main

import (
	db "TutorialGoLang/35_package-initialization/database"
	_ "errors"
	"fmt"
)

func main() {
	database := db.GetConnection()

	fmt.Println("Database connection is", database)
}
