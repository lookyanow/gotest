package main 

//import "fmt"
import "net/http"

func main(){
	http.HandleFunc("/", someFunc)
	http.ListenAndServe(":8081", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello world2\n"))
}
