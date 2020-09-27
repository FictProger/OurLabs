package main
import (
	"log"
	"net/http"
	
)
​
func main() {
​

​
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {}
	
	log.Println("Staring HTTP server...")
	log.Fatal(http.ListenAndServe(":8795", nil))
}