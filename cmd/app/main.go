package main

import (
	"database/sql"
	"github.com/whyistilley/docker-golang/internal/database"
	"github.com/whyistilley/docker-golang/internal/service"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	pqConnStr string
	svc       *service.Service
	db        *sql.DB
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

func init() {
	var err error

	svc = service.New("app")

	pqConnStr = os.Getenv("POSTGRES_URI")
	db, err = database.New("postgres", pqConnStr)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	defer db.Close()
	log.Println(svc.BuildInfo())

	Ping()

	users := Users()
	log.Println("Users():", users)

	userByEmail := UserByEmail("charlie@example.com")
	log.Println("UserByEmail():", userByEmail)

	newUser := InsertUser("Agatha", "aharkness@westview.com")
	log.Println("InsertUser():", newUser)
	userById := UserById(newUser.Id)
	log.Println("UserById():", userById)

	users = Users()
	log.Println("Users():", users)

	userByEmail = UserByEmail("aharkness@westview.com")
	log.Println("UserByEmail():", userByEmail)

	deletedUser := DeleteUserByEmail("aharkness@westview.com")
	log.Println("DeleteUserByEmail():", deletedUser)

	updatedUser := UpdateUserNameByEmail("alice@example.com", "Allison")
	log.Println("UpdateUserNameByEmail():", updatedUser)

	userByEmail = UserByEmail("alice@example.com")
	log.Println("UserByEmail():", userByEmail)

	updatedUser = UpdateUserNameById(userByEmail.Id, "Ali")
	log.Println("UpdateUserNameById():", updatedUser)
}

func Ping() {
	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connection established...")
}

func Users() []User {
	statement, err := db.Prepare("SELECT * FROM get_users()")
	if err != nil {
		log.Fatalln(err)
	}
	rows, err := statement.Query()
	if err != nil {
		log.Fatalln(err)
	}

	var user User
	var users []User

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt); err == nil {
			users = append(users, user)
		}
	}
	return users
}

func UserByEmail(email string) *User {
	var user User
	statement, err := db.Prepare("SELECT * FROM get_user_by_email($1)")
	if err != nil {
		log.Println(err)
	}
	defer statement.Close()
	row := statement.QueryRow(email)
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return &user
}

func UserById(id string) *User {
	var user User
	//statement, err := db.Prepare("SELECT * FROM users WHERE id = $1 LIMIT 1")
	statement, err := db.Prepare("SELECT * FROM get_user_by_id($1)")
	if err != nil {
		log.Println(err)
	}
	defer statement.Close()
	row := statement.QueryRow(id)
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return &user
}

func InsertUser(name, email string) *User {
	statement, err := db.Prepare("SELECT * FROM insert_user($1, $2)")
	if err != nil {
		log.Println(err)
	}
	defer statement.Close()
	var newUser User
	row := statement.QueryRow(name, email)
	err = row.Scan(&newUser.Id, &newUser.Name, &newUser.Email, &newUser.CreatedAt)
	if err != nil {
		log.Println(err)
		return &User{}
	}
	return &newUser
}

func DeleteUserByEmail(email string) *User {
	statement, err := db.Prepare("SELECT * FROM delete_user_by_email($1)")
	if err != nil {
		log.Println(err)
	}
	defer statement.Close()
	row := statement.QueryRow(email)
	var deletedUser User
	err = row.Scan(&deletedUser.Id, &deletedUser.Name, &deletedUser.Email, &deletedUser.CreatedAt)
	if err != nil {
		log.Println(err)
		return &User{}
	}
	return &deletedUser
}

func UpdateUserNameByEmail(email, name string) *User {
	statement, err := db.Prepare("SELECT * FROM update_user_by_email($1, $2)")
	if err != nil {
		log.Println(err)
	}
	defer statement.Close()
	row := statement.QueryRow(email, name)
	var updatedUser User
	err = row.Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Email, &updatedUser.CreatedAt)
	if err != nil {
		log.Println(err)
		return &User{}
	}
	return &updatedUser
}

func UpdateUserNameById(id, name string) *User {
	statement, err := db.Prepare("SELECT * FROM update_user_by_id($1, $2)")
	if err != nil {
		log.Println(err)
	}
	defer statement.Close()
	row := statement.QueryRow(id, name)
	var updatedUser User
	err = row.Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Email, &updatedUser.CreatedAt)
	if err != nil {
		return &User{}
	}
	return &updatedUser
}
