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

	fmt.Println("connected......")
	// insert data into databae 
	insertDB := `insert into "friends" ("name", "number") values ('mygee','+000333222')`
	// execute database insert call 
	 _ , e := db.Exec(insertDB)
	//check error
	checkError(e)
	fmt.Println("Success ..... ")
}
func main(){
	conn_database()
}
/* to delete use the -- `delete from "{dbname}"  where {your choice}`
	then execute your query..

	to update ---- ` update {dbname} set {xxx,xxx} where {your choice}`

*/

