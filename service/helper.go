package service 
import (
	"database/sql"
	"log"

)
//struct for Emmloyee
type Employee struct{

ID int `json:"id"`
Name string `json:"name"`
City string `json:"city"`
PinCode int `json:"pincode"`	

}

func FetchAllEmployee(db *sql.DB)[]Employee{
log.Println("Fetching Employee....");
result, err:=db.Query("SELECT * FROM Employee ORDER BY id DESC");
if err != nil {
	panic(err.Error())
}
emp := Employee{}
res := []Employee{}

for result.Next() {
	var id , pincode int
	var name, city string
	err = result.Scan(&id, &name, &city, &pincode)
	if err != nil {
		panic(err.Error())
	}
	emp.ID = id
	emp.Name = name
	emp.City = city
	emp.PinCode= pincode
	res = append(res, emp)
}
defer db.Close();
return res


}