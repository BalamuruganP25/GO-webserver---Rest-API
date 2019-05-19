package main 

import(

   "fmt"
   "net/http"
   "encoding/json"
   "bytes"
   "io/ioutil"
)

 type login struct{

 	Username string `json:"username"`
 	Password string `json:"password"`
 }


 type loginresponse struct{

  Loginstatus string `json:loginstatus`

 }
 

func main(){


    var getuserlogindetails login

    getuserlogindetails.Username = "balamurugan.p@XXXX.in" 
    getuserlogindetails.Password = "abcd1234"

    convert_marshal_logindata, _ := json.Marshal(getuserlogindetails)
    fmt.Println("Send request")

    req, _ := http.NewRequest("POST","http://localhost:8088/login", bytes.NewBuffer(convert_marshal_logindata))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()

}