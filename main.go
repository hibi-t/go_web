package main

import (
	//"database/sql"
	//"log"
	//"models"
	"os"
	//"github.com/vattle/sqlboiler/boil"
)

func main() {
    DB := os.Getenv("DB")
    println(DB)

    //db, err := sql.Open(DB, "dbname="+DBNAME+ "user="+DBUSER+ "password="+DBPW+ "sslmode="+SSLMODE)
//     db, err := sql.Open("postgres", "dbname=DBNAME user=DBUSERNAME sslmode=disable")
//     if err != nil {
//         log.Println(err)
//         return
//     }
//
//     boil.SetDB(db)
//
//     users, err := models.UsersG().All()
//     if err != nil {
//         log.Println(err)
//         return
//     }
//
//     pp.Println(users)
}
