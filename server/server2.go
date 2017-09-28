package main 

import "net/http"

func main(){

	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8085", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Push triggered Jenkins build test2\n"))
}

