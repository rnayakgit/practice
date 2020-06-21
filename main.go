package main
import (
	"encoding/json"
	"fmt"
	"time"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"controller"
)
func init(){
	fmt.Println("INIT Executed .....");
}

func healthCheck(w http.ResponseWriter, r *http.Request){
timedetails:=time.Now().Local().String();
log.Println("healthCheck Api executed....");
json.NewEncoder(w).Encode(map[string]string{"Time":string(timedetails),
    "status": "Ok" });


}


func main (){
fmt.Println(" *********** Main Executed *************");
router:=mux.NewRouter();
router.Headers().HeadersRegexp("Content-Type", "application/json");

//api's
router.HandleFunc("/health",healthCheck);
router.HandleFunc("/getEmployee",controller.GetEmployee);

//router server 
log.Fatal(http.ListenAndServe(":9999", router));
fmt.Println(" *********** Main Ended *************");
}