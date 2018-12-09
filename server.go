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

func creating(w http.ResponseWriter, r * http.Request)  {
	println("linked for creating.")
	var name = r.FormValue("name")
	var psw = r.FormValue("psword")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")

	cfg.DBName = "users"
	db, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot get the database")
		return
	}

	var q = "select password from wholepeople where username='" + name + "'"
	jpj, err := db.Query(q)

	if err != nil {
		w.Write([]byte("-1"))
		println("connecting pro. "  + err.Error())
		return
	}

	var dbName string
	jpj.Next()
	jpj.Scan(&dbName)

	if dbName != "" {
		println("dbname is: ... " + dbName )
		w.Write([]byte("shabi"))
		return
	}

	q = "INSERT INTO wholepeople (username, password) VALUES (\"" + name + "\", \"" + psw + "\")"


	_, err = db.Query(q)

	if err == nil{
		println("succed.")
		w.Write([]byte("1"))
		return
	}

	println("shabi....")
	println(err.Error())
	w.Write([]byte("-1"))
}

func signing(w http.ResponseWriter, r * http.Request)  {
	println("linked for signing.")
	var name = r.FormValue("name")
	var psw = r.FormValue("psword")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")

	cfg.DBName = "users"
	db, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot get the database")
		return
	}

	var q = "select password from wholepeople where username = '" + name + "'"
	pasw, err := db.Query(q)

	if err != nil {
		println(err.Error())
		println("something wrong with username or passworld")
		w.Write([]byte("-1"))
		return
	}
	var dbName string
	pasw.Next()
	pasw.Scan(&dbName)

	if psw == dbName {
		println("right user is coming in")
		w.Write([]byte("1"))
	}else {
		println(dbName + " is not a correct psw")
		w.Write([]byte("shabi, hehe, -1"))
	}





/* it seems that we do in the users database ~~~ happy!
	table, err := db.Query("SHOW TABLES")
	if err != nil {
		w.Write([]byte("sorry, cannot"))
		println("we cannot get in users")
		return
	}
	var dbName string
	table.Next()
	err = table.Scan(&dbName)
	if err != nil{
		println("error for table" + err.Error())
		w.Write([]byte("error"))
	}
	println(dbName)
	w.Write([]byte(dbName))

test finished */




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
	http.HandleFunc("/create", creating)

	//var name  = "jojo"
	//var psw = "heh"
	//println("INSERT INTO wholepeople (username, password) VALUES (\"" + name + "\", \"" + psw + "\")")

	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil)


}
// [END main]
