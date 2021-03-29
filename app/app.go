package app

import "net/http"

func Start()  {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe("localhost:8000", mux)
}
