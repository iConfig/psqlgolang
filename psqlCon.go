package main

import (
	"fmt"
	"database/sql"  // importing sql module
	_ "github.com/lib/pq" // importing psql module , don't forget the '_' sign
	// might need to run "go get github.com/lib/pq"
	
)
// declaring connection constant
const(
	host = "localhost"
	port = 5432
	username = "postgres"
	password = "postgres"
	database_name = "mydata"
)
// default error check function
func checkError(err error){
	if err !=nil{
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

	fmt.Println("connected......")
}

func main(){
	conn_database()
}