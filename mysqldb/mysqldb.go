package mysqldb
import (
    _ "github.com/go-sql-driver/mysql"
"database/sql"
"log"
)

//dbconnetion
func DbConn() (db *sql.DB) {
  //  dbDriver := "mysql"
    //dbUser := "root"
    //dbPass := "root"
    //dbName := "test"
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
    if err != nil {
		log.Println("error occurred in conection:");
        panic(err.Error())
	}
    return db
}