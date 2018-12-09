package main

import (
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
	//var name = r.FormValue("name")
	//var psw = r.FormValue("psword")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")

	cfg.DBName = "users"
	db, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot get the database")
		return
	}

	table, err := db.Query("SHOW TABLES")

	if err != nil {
		w.Write([]byte("sorry, cannot"))
		println("we cannot get in users")
		return
	}

	var dbName string
	table.Scan(dbName)
	w.Write([]byte(dbName))

	//rows, err := db.Query("SHOW DATABASES")
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Could not query db: %v", err), 500)
	//	return
	//}
	//defer rows.Close()
	//
	//buf := bytes.NewBufferString("Databases:\n")
	//for rows.Next() {
	//	var dbName string
	//	if err := rows.Scan(&dbName); err != nil {
	//		http.Error(w, fmt.Sprintf("Could not scan result: %v", err), 500)
	//		return
	//	}
	//	fmt.Fprintf(buf, "- %s\n", dbName)
	//}
	//w.Write(buf.Bytes())



	//w.Write([]byte("linked to the fxcking database..."))
	//println(name, psw)
	//
	//
	//println("cnm...")
	//w.Write([]byte("linked..."))
	//db, err := sql.Open("firstnote", "root:root@/uestcbook")
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
