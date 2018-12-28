package main 

//import "fmt"
import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main(){
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/dump", getDump)
	http.ListenAndServe(":80", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Index responce from server\n"))
}

func hello(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Hello world\n"))
}

func getDump(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		fmt.Fprint(w, string(requestDump))
		log.Print(string(requestDump))
	}
}