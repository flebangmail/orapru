package db

import (
	"database/sql"
	"fmt"
	"time"
)

var localDB = map[string]string{
	"service":  "XE",
	"username": "juarez",
	"server":   "190.210.215.185",
	"port":     "1521",
	"password": "jrz", 

func ChequeoConexion() bool {

	fmt.Println("*** Using only go_ora package (no additional client software)")
	fmt.Println("Local Database, simple connect string ")
	t := time.Now()
	doDBThings(localDB)
	fmt.Println("Time Elapsed", time.Since(t).Milliseconds())

	return true
}

func doDBThings(dbParams map[string]string) {
	connectionString := "oracle://" + dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"]
	//if val, ok := dbParams["walletLocation"]; ok && val != "" {
	//	connectionString += "?TRACE FILE=trace.log&SSL=enable&SSL Verify=false&WALLET=" + url.QueryEscape(dbParams["walletLocation"])
	//}
	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println("Can't close connection: ", err)
		}
	}()

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

}
