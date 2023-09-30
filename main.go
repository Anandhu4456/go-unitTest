package main

import "net/http"

func main() {
	http.HandleFunc("/",DemoHandler)
	http.ListenAndServe(":3000",nil)
}

func DemoHandler(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hey"))
}
