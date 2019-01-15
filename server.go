package main 

//import "fmt"
import (
	"context"
	"fmt"
	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
)

func newRequestId() string{
	id := xid.New()
	return id.String()
}

type myKey string

var requestIDKey = myKey("requestID")
func main(){
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/dump", getDump)
	http.HandleFunc("/live", liveCheck)
	http.HandleFunc("/ready", readyCheck)
	http.ListenAndServe(":80", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("Index responce from server\n"))
}

func hello_call(ctx context.Context, str string) string{
	log.WithField("request_id", ctx.Value(requestIDKey)).Info("hello_call function called")
	return str + "s"
}

func hello(w http.ResponseWriter, req *http.Request){
	ctx := req.Context()
	ctx = context.WithValue(ctx, requestIDKey, newRequestId())
	log.WithField("request_id", ctx.Value(requestIDKey)).Info("hello handler called")
	str := hello_call(ctx, "hello")
	log.WithField("request_id", ctx.Value(requestIDKey)).Info("hello handler after strings", str)
	w.Write([]byte("Hello world\n"))
}

func liveCheck(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func readyCheck(w http.ResponseWriter, req *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ready"))
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

