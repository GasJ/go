package main

import (
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"net/http"
	"time"

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

	q = "INSERT INTO wholepeople (username, password, imageid) VALUES (\"" + name + "\", \"" + psw + "\", -1)"


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
	defer db.Close()


	var q = "select password from wholepeople where username = '" + name + "'"
	pasw, err := db.Query(q)
	defer pasw.Close()


	if err != nil {
		println(err.Error())
		println("something wrong with username")
		w.Write([]byte("-1"))
		return
	}
	var dbpsw string
	pasw.Next()
	err = pasw.Scan(&dbpsw)
	if err != nil{
		println(dbpsw + err.Error())
	}


	if psw == dbpsw {
		println("right user is coming in")

		db, err := mysql.DialCfg(cfg)


		//q = "select imageid from wholepeople where username = '" + name + "';"
		wwwww := "select password from wholepeople where username = '" + name + "'"
		println(q)
		image, err := db.Query(wwwww)
		jojo, err := db.Query(q)
		defer image.Close()


		if err != nil{
			println("cnm, the problem i workd for hours is: " + err.Error())
		}

		var id string
		jojo.Next()
		for image.Next(){
			err := image.Scan(id)
			if err != nil{
				println("cnm shab " + err.Error())
			}
			println(id)
		}

		switch id {
		case "1":
			w.Write([]byte("11"))
			println("user go with an image")
		case "2":
			w.Write([]byte("12"))
			println("user go with an image")
		case "3":
			w.Write([]byte("13"))
			println("user go with an image")
		case "4":
			w.Write([]byte("14"))
			println("user go with an image")
		case "5":
			w.Write([]byte("15"))
			println("user go with an image")
		case "6":
			w.Write([]byte("16"))
			println("user go with an image")
		case "7":
			w.Write([]byte("17"))
			println("user go with an image")
		case "8":
			w.Write([]byte("18"))
			println("user go with an image")
		case "9":
			w.Write([]byte("19"))
			println("user go with an image")
		case "10":
			w.Write([]byte("110"))
			println("user go with an image")
		default:
			w.Write([]byte("111"))
			println("user go without an image")
			println(id)
		}

	}else {
		println(dbpsw + " is not a correct psw")
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

func welc(w http.ResponseWriter, r * http.Request)  {
	println("linked for welcome.")
	var name = r.FormValue("name")
	var img = r.FormValue("img")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")

	cfg.DBName = "users"
	db, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot get the database")
		return
	}

	var q = "UPDATE wholepeople SET imageid=\"" + img + "\" WHERE username='" + name + "'"
	_, err = db.Query(q)

	if err != nil {
		w.Write([]byte("-1"))
		println("cannot set imageid. "  + err.Error())
		return
	}

	w.Write([]byte("hehe"))
	println(name + "go with an imageid.")

}

func createplan(w http.ResponseWriter, r * http.Request){
	println("linked for welcome.")
	var name = r.FormValue("name")
	var planname = r.FormValue("planname")
	var imp = r.FormValue("importance")
	var disc = r.FormValue("discription")
	var reminder = r.FormValue("remind")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")

	cfg.DBName = "users"
	db, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot get the database")
		return
	}

	var q = "INSERT INTO plan (user, createdate，planname, import) VALUES ( \"" + name + "\", " +
		time.Now().Format("2018-12-10 23:56:30") + ",\"" + planname + "\"," + imp + ")"
	_, err = db.Query(q)

	if err != nil{
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot create a plan")
		return
	}

	if disc != ""{
		q = "update plan set discription = \"" + disc + "\" where user=\"" +
			name + "\" and planname=\"" + planname + "\""
		_, err = db.Query(q)
		if err != nil{
			println(err.Error())
			println("we cannot add disc")
		}
	}

	if reminder != ""{
		q = "update plan set timeremindend= \"" + reminder + "\" where user=\"" +
			name + "\" and planname=\"" + planname + "\""
		_, err = db.Query(q)
		if err != nil{
			println(err.Error())
			println("we cannot add reminder")
		}
	}

	w.Write([]byte("1"))
}

func getwhole(w http.ResponseWriter, r * http.Request){

}

// [END handler]

// [START main]
func main() {
	http.HandleFunc("/stickynote", handler)
	http.HandleFunc("/signin", signing)
	http.HandleFunc("/create", creating)
	http.HandleFunc("/welcome", welc)
	http.HandleFunc("/cards/createplan", createplan)
	http.HandleFunc("/cards/getwholeplan", getwhole)

	start := time.Now()
	print(start.Format("2006-01-02 15:04:05"))

	//var name  = "jojo"
	//var psw = "heh"
	//println("INSERT INTO wholepeople (username, password) VALUES (\"" + name + "\", \"" + psw + "\")")

	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil)


}
// [END main]
