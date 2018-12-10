package main

import (
	"database/sql"
	"fmt"
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

	if dbName == psw {
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



	var q = "select password, imageid from wholepeople where username = '" + name + "'"
	pasw, err := db.Query(q)
	defer pasw.Close()


	if err != nil {
		println(err.Error())
		println("something wrong with username")
		w.Write([]byte("-1"))
		return
	}
	var dbpsw string
	var id string
	pasw.Next()
	err = pasw.Scan(&dbpsw, &id)
	if err != nil{
		println("shabi " + err.Error())
	}

	if psw == dbpsw {
		println("right user is coming in")

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
	println("linked for creating plan.")
	var name = r.FormValue("name")
	var planname = r.FormValue("planname")
	var imp = r.FormValue("importance")
	var disc = r.FormValue("discription")
	var reminder = r.FormValue("remind")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")

	cfg.DBName = "users"
	db, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("-3"))
		println(err.Error())
		println("we cannot get the database")
		return
	}



	var biubiu = "SELECT user from plan Where planname = '" + planname + "'"
	jpj, err := db.Query(biubiu)

	if err != nil{
		w.Write([]byte("-2"))
		println("cannot get...")
		return
	}


	var dbName string
	var jo = true
	for jo{
		jpj.Next()
		err = jpj.Scan(&dbName)
		println("name: " + name + "curname: " + dbName)
		if dbName == name {
			println("it exists, and it is: ... " + dbName + name )
			w.Write([]byte("-1"))
			return
		} else if dbName == "" || err != nil {
			jo = false
		}
	}



	var q = "INSERT INTO plan (user, planname, import) VALUES ( \"" + name + "\", \"" + planname + "\"," + imp + ")"
	println(q)
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
	println("linked for whole plans.")
	var name = r.FormValue("name")

	//cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")
	//
	//cfg.DBName = "users"
	//db, err := mysql.DialCfg(cfg)
	//
	//if err != nil {
	//	w.Write([]byte("-3"))
	//	println(err.Error())
	//	println("we cannot get the database")
	//	return
	//}


	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "starvingmonkey", "lyzsb", "104.154.216.44", "users" )
	db, err := sql.Open("mysql", dsn)


	var q = "select planname, import, discription from plan WHERE user='" + name + "'"
	println(q)
	rows, err := db.Query(q)



	if err != nil{
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot get your plan")
		return
	}

	aaa, _ :=rows.Columns()
	println("column size: " , len(aaa))


	var (
		id        string
		firstName string
		lastName  string

	)
	var msg []byte
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			println(err.Error())
		}
		println("%v: %s %s", id, firstName, lastName)
		mmm := append(msg, []byte(id + "," + string(firstName) + "," + lastName + "\n")...)
		msg = mmm
	}
	err = rows.Err()
	if err != nil {
		println(err.Error())
	}

	w.Write(msg)

	//var pn string
	//var ip string
	//var di string
	//var jo = true
	//var msg []byte
	//for jo{
	//	//jpj.Next()
	//	jpj.NextResultSet()
	//	err = jpj.Scan(&pn, &ip, &di)
	//	if err != nil{
	//		println("shabi " + err.Error())
	//		//jo = false
	//		if err.Error() == "sql: Rows are closed"{
	//			jo=false
	//		}
	//	}
	//	if pn == ""{
	//		jo = false
	//	}
	//	println("current: " + pn + ip + di)
	//	mmmm := append([]byte(pn + "," + string(ip) + "," + di + "\n"))
	//	msg = mmmm
	//}





}

func editplans(w http.ResponseWriter, r * http.Request){
	println("linked for creating plan.")
	var name = r.FormValue("name")
	var planname = r.FormValue("planname")
	var imp = r.FormValue("importance")
	var disc = r.FormValue("discription")
	var oldone = r.FormValue("oldone")

	//cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")
	//
	//cfg.DBName = "users"
	//db, err := mysql.DialCfg(cfg)
	//
	//if err != nil {
	//	w.Write([]byte("-3"))
	//	println(err.Error())
	//	println("we cannot get the database")
	//	return
	//}
	//
	//
	//var q = "UPDATE table plan SET planname='" + planname +"', import="+ imp +"   WHERE user=\"" +
	//	name + "\" and planname=\"" + oldone + "\""
	//println(q)
	//_, err = db.Query(q)
	//
	//if err != nil{
	//	w.Write([]byte("-1"))
	//	println(err.Error())
	//	println("we cannot eidt a plan")
	//	return
	//}
	//
	//if disc != ""{
	//	q = "update plan set discription = \"" + disc + "\" where user=\"" +
	//		name + "\" and planname=\"" + planname + "\""
	//	_, err = db.Query(q)
	//	if err != nil{
	//		println(err.Error())
	//		println("we cannot add disc")
	//	}
	//}


	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "starvingmonkey", "lyzsb", "104.154.216.44", "users" )
	db, err := sql.Open("mysql", dsn)


	var q = "select planname, import, discription from plan WHERE user='" + name + "'"
	println(q)
	rows, err := db.Query(q)



	if err != nil{
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot get your plan")
		return
	}

	aaa, _ :=rows.Columns()
	println("column size: " , len(aaa))


	var (
		id        string
		firstName string
		lastName  string

	)
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			println(err.Error())
		}
		println("%v: %s %s", id, firstName, lastName)
		if id == oldone{
			var q = "UPDATE table plan SET planname='" + planname +"', import="+ imp +"   WHERE user=\"" +
				name + "\" and planname=\"" + oldone + "\""
			println(q)
			_, err = db.Query(q)
			if err != nil{
				println(err.Error())
				println("we cannot edit it")
			}
			if disc != ""{
				q = "update plan set discription = \"" + disc + "\" where user=\"" +
					name + "\" and planname=\"" + planname + "\""
				_, err = db.Query(q)
				if err != nil{
					println(err.Error())
					println("we cannot add disc")
				}
				w.Write([]byte("1"))
		}}
	}
	err = rows.Err()
	if err != nil {
		println(err.Error())
	}

		w.Write([]byte("-1"))


}

func rmplan(w http.ResponseWriter, r * http.Request){
	println("linked for removing plan.")
	var name = r.FormValue("name")
	var planname = r.FormValue("planname")

	cfg := mysql.Cfg("glossy-radio-224901:us-central1:firstnote", "starvingmonkey", "lyzsb")

	cfg.DBName = "users"
	db, err := mysql.DialCfg(cfg)

	if err != nil {
		w.Write([]byte("-3"))
		println(err.Error())
		println("we cannot get the database")
		return
	}


	var q = "delete from plan WHERE planname='" + planname +"' and user = '" + name +"'"
	println(q)
	_, err = db.Query(q)

	if err != nil{
		w.Write([]byte("-1"))
		println(err.Error())
		println("we cannot create a plan")
		return
	}

	w.Write([]byte("1"))

}

// [END handler]

// [START main]
func main() {
	http.HandleFunc("/stickynote", handler)
	http.HandleFunc("/signin", signing)
	http.HandleFunc("/create", creating)
	http.HandleFunc("/welcome", welc)
	http.HandleFunc("/cardscreate", createplan)
	http.HandleFunc("/cards/getwholeplan", getwhole)
	http.HandleFunc("/cardsedit", editplans)
	http.HandleFunc("/cardsremove", rmplan)

	start := time.Now()
	print(start.Format("2006-01-02 15:04:05"))

	//var name  = "jojo"
	//var psw = "heh"
	//println("INSERT INTO wholepeople (username, password) VALUES (\"" + name + "\", \"" + psw + "\")")

	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS(":443", "ssl.crt", "ssl.key", nil)


}
// [END main]
