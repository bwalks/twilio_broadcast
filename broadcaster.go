package broadcaster

import (
	"fmt"
	"net/http"
)

func init(){
	http.HandleFunc("/", test)
	http.HandleFunc("/twilio/broadcast", broadcast)
	http.ListenAndServe(":8000", nil)
}


func test(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}

func broadcast(w http.ResponseWriter, r *http.Request) { 
	fmt.Fprint(w, "Broadcasting")
}