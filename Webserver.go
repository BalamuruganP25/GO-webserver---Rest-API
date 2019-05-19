package main 

import(

   "fmt"
   "net/http"
   "github.com/gorilla/mux"
   "io/ioutil"
   "encoding/json"
   "strings"
)


type login struct{

Username string `json:"username"`
Password string `json:"password"`

}


type loginresponse struct{

 Loginstatus string `json:loginstatus`

}

func userlogin(w http.ResponseWriter, r *http.Request){

	var resultlogin loginresponse
    fmt.Println(">>>>>>> User Login <<<<<<<<<<")
    var recevielogindata login
    jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading the body", err)
	}
	err = json.Unmarshal(jsn, &recevielogindata)
	if err != nil {
		fmt.Println("Decoding error: ", err)
	}
   fmt.Println("Received: %v\n", recevielogindata)

   if strings.TrimRight(recevielogindata.Username, "\n") == "balamurugan.p@XXX.in" {
  
  	resultlogin.Loginstatus =  "success"
   }else{


   	resultlogin.Loginstatus =  "falied"


   }
  
   reposnseforlogin, err := json.Marshal(resultlogin)
   if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(reposnseforlogin)
}

func main(){

incomingrequest := mux.NewRouter()
incomingrequest.HandleFunc("/login",userlogin).Methods("POST")
http.ListenAndServe(":8088", incomingrequest)

}


