package main 

//import "fmt"
import "net/http"

func main(){
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8888", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Index responce from server\n"))
}

func hello(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello world\n"))
}
