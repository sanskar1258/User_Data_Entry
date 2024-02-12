package models

import (
	"fmt"

	"dataentry/utility"

	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(name, address string) {

	fmt.Println("Connection Established successfully!")

	var err error

	stmt, err := utility.Db.Prepare("Insert into users(name, address) values (?,?)")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, address)

	if err != nil {
		panic(err)
	}

	fmt.Println("User inserted successfully!")

}
