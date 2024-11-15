package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/whyistilley/docker-golang/internal/service"
)

var (
	pqConnStr string
)

func init() {
	pqConnStr = os.Getenv("POSTGRES_URI")
}

func main() {
	app := service.New("app")
	log.Println(app.BuildInfo())
	db, err := sql.Open("postgres", pqConnStr)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connection established...")
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatalln(err)
	}
	var user struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	var users []struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}
	log.Println(users)
}
