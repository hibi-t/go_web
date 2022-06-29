package main

import (
	"database/sql"
	//"log"
	//"models"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	// models "go_web/models"
	_ "github.com/lib/pq"
)

// id bigserial NOT NULL,
// title varchar(255) NOT NULL,
// note text,
// finished boolean NOT NULL,
// due_date timestamp,
// CONSTRAINT textnote_pkc PRIMARY KEY(id)

func initDB() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Printf("読み込み失敗\n")
	}
	env := os.Getenv("ENV")
	db := os.Getenv("DB")
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")

	dns := "host="+env+" dbname="+dbname+" user="+dbuser+" password="+dbpass+" sslmode=disable connect_timeout=10"
	fmt.Println("db:", dns, db)

	con, err := sql.Open(db, dns)
	if err != nil {
		panic(err)
	}
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(10)
	con.SetConnMaxLifetime(300 * time.Second)

	boil.SetDB(con)
	boil.DebugMode = true
}

// func insert() {
// 	textnote := db.Textnote{title: null.StringFrom("タイトル"), note: null.StringFrom("記事の内容"), finished: null.StringFrom(true), due_date: null.StringFrom("2022-06-29")}
// 	fmt.Printf("before textnote = %+v\n", textnote)
// 	user.InsertGP(context.Background(), boil.Infer())
// 	fmt.Printf("after textnote = %+v\n", textnote)
// }

func main() {
	initDB()
	// insert()
}
