package main

import (
	//"database/sql"
	//"log"
	//"models"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Printf("読み込み失敗\n")
	}
	env := os.Getenv("ENV")
	DB := os.Getenv("DB")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASS")

	fmt.Println(env)
	fmt.Println(DB)
	fmt.Println(DBUser)
	fmt.Println(DBPass)
}
