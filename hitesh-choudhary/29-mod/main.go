/*


 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in Golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc( "/", serveHome).Methods("GET")
	http.ListenAndServe(":4000", r)
	log.Fatal()	
}

func greeter(){
	fmt.Println("Hey there mod users")
}

 // w is http response writer  and r is a type of pointer handling request
func serveHome( w http.ResponseWriter, r *http.Request){
	w.Write([]byte("	<h1>Welcome to golang series on Youtube</h1>"))
}