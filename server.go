package main
​
import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)
​
func main() {
​
	type Result struct {
		Time string `json:"time"`
	}
​
	http.HandleFunc("/time", func(rw http.ResponseWriter, r *http.Request) {
		res := Result{time.Now().Format(time.RFC3339)}
​
		rw.Header().Set("content-type", "application/json")
		rw.WriteHeader(200)
		_ = json.NewEncoder(rw).Encode(res)
	})
​
	log.Println("Staring HTTP server...")
	log.Fatal(http.ListenAndServe(":8795", nil))
}