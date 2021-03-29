package app

import (
"encoding/json"
"fmt"
_ "fmt"
"net/http"
_ "net/http"
)

type Customer struct {
	Name string
	City string
	Zipcode string
}

func greet (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request)  {
	customers := []Customer {
		{"Bi", "HCM", "11111"},
		{"Bi2", "HCM", "22222"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}




