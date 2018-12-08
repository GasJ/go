package main

import (
	//"database/sql"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
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
	println("linked.")
	var name = r.FormValue("name")
	var psw = r.FormValue("psword")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "a")
	cfg.DBName = "users"
	_, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("sha bi"))
		return
	}


	w.Write([]byte("linked to the fxcking database..."))
	println(name, psw)



	//db, err := sql.Open("firstnote", "root:root@/uestcbook")
	//
	//result, err := db.Exec(
	//	"INSERT INTO users (name, age) VALUES (?, ?)",
	//	"gopher",
	//	27,
	//)
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
