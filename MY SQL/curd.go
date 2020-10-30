package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "kloudlearn"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	fmt.Println("Go connecting to mysql")

	db, err := sql.Open("mysql", dsn(dbname))

	if err != nil {
		panic(err.Error())
		return
	}
	fmt.Println("Connected")
	defer db.Close()

	insert, err := db.Query("INSERT INTO user_info (name,mobile,address) VALUES ('SUMIT','8948985801','MM/349')")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully inserted the data")

	defer insert.Close()
	update, err := db.Query("Update user_info SET name='Rohan', mobile ='123456742' where id = 2")

	if err != nil {
		panic(err.Error())
	}

	defer update.Close()

	delete, err := db.Query("DELETE FROM user_info WHERE id = 3")

	if err != nil {
		panic(err.Error())
	}

	defer delete.Close()

	results, err := db.Query("Select * from user_info")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var id int
		var name string
		var address string
		var mobile string

		err = results.Scan(&id, &name, &address, &mobile)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(address)
		fmt.Println(mobile)
	}
	defer results.Close()

}
