package main

import (
	"fmt"
	"database/sql" //importing sql module 
	_ "github.com/lib/pq" // importing psql module
)
// declaring connection constant
const (
	host = "localhost"
	port = 5432
	username = "postgres"
	password = "postgres"
	database_name = "mydata"
)
// default error check function 
func checkError(err error){
	if err != nil {
		panic(err)
	}
}
// database connection function
func conn_database(){
	// database connection string  
	psqlcon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
				host,port,username,password,database_name)
	// open database
	db , err := sql.Open("postgres", psqlcon)
	// check for error 
	checkError(err)
	//close database on function return 
	defer db.Close()
	//check database 
	err = db.Ping()
	checkError(err)
	fmt.Println("connected......\n")
	// select data from databae 
	row, err := db.Query(`select "name" , "number" from "friends"`)
	//check error
	checkError(err)
	//close database on function return 
	defer row.Close()
	// iterating over the rows 
	for row.Next(){
		//declaring string var for both name and number 
		var name string
		var number string
		//reading columns in each rows into variables
		err = row.Scan(&name, &number)
		checkError(err)

		fmt.Println(name, number)
	}
	fmt.Println("\nSuccess ..... ")
}

func main(){
	conn_database()
}

