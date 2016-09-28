package main

import "fmt"
import "os"
import "net/http"
import "io/ioutil"
import "time"

func main(){
	url := os.Args[1]
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil{
		fmt.Printf("Error %v\n",err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("StatusCode: %d \n",resp.StatusCode)
	fmt.Printf("Responce time: %.2fs\n", time.Since(start).Seconds())
	fmt.Printf("Content: %s\n",b)
	
	
}