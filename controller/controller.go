package controller

import (
	"log"
	"net/http"
	"encoding/json"
	"mysqldb"
	"service"

)

func GetEmployee(w http.ResponseWriter, r *http.Request){
log.Println("*****getEmployee Api called************8");
conn:=mysqldb.DbConn();
emp:=service.FetchAllEmployee(conn);
json.NewEncoder(w).Encode(emp);
}
