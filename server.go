package main

import (
	"net/http"
)

// [START handler]
func handler(w http.ResponseWriter, r * http.Request) {
	println("linked.")
	var data = r.FormValue("message")
	if data == "we"{
		println("hey, sb.")
	}
	var psw = r.FormValue("message2")
	if psw == "ohoh" {
		println("hahahahahaha")
	}
	w.Write([]byte("hey, friend."))
}

func signing(w http.ResponseWriter, r * http.Request)  {
	
}


// [END handler]

// [START main]
func main() {
	http.HandleFunc("/stickynote", handler)
	http.HandleFunc("/signin", signing)
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil)
}
// [END main]
