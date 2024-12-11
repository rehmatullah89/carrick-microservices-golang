package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//connecting database
	db, err := sql.Open("mysql", "rehmatullah:Allahi$great1@tcp(127.0.0.1:3306)/sakila")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//creating table through SQL DDL commands
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sakila.user(id int, name varchar(50), email varchar(50))")
	if err != nil {
		panic(err.Error())
	}

	//inserting 5 records
	result, err := db.Exec("INSERT INTO sakila.user(id,name,email) VALUES (101,'rehmat','rehmat@gmail.com'),(102,'Shahbaz','shahbaz@gmail.com'),(103,'Tasawar','tswar@gmail.com'),(104,'mashood','mashood@gmail.com'),(105,'faisal','faisal@gmail.com')")

	if err != nil {
		panic(err.Error())
	}

	rc, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("inserted %d rows\n", rc)

	//simple select query to fetch all records

	rows, err := db.Query("SELECT * FROM sakila.user")
	if err != nil {
		panic(err.Error())
	}

	// use this struct to populate results returned from the database
	type user struct {
		ID    int
		Name  string
		Email string
	}

	for rows.Next() {
		var e user
		err = rows.Scan(&e.ID, &e.Name, &e.Email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d %s %s\n", e.ID, e.Name, e.Email)
	}

}
