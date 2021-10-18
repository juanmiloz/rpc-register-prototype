package main

//import "errors"
import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string `json:"username"`
	Email        	string `json:"email"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Password        string `json:"password"`
	Country        	string `json:"country"`
}

var user_logged User

func createNewUser(user User) (error, int) {
	response := checkUsernameCreate(user)
	if(response == 1) {
		log.Println("Username repeated")
		return nil, 2	
	}
	response = checkEmailCreate(user)
	if(response == 1) {
		log.Println("Email repeated")
		return nil, 3	
	}
	q :=`INSERT INTO users
		(username, email, firstname, lastname, password, country) 
		VALUES ($1,$2,$3,$4,$5,$6)`
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil{
		return err, 0
	}
	defer stmt.Close()

	password_byte := []byte(user.Password)
	bcrypt_password, err := bcrypt.GenerateFromPassword(password_byte, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    log.Println(string(bcrypt_password))

	_, err = stmt.Exec(user.Username, user.Email, user.Firstname, user.Lastname, bcrypt_password, user.Country)
	if err != nil{
		return err, 0
	}

	_ = loginProcess(user.Email, user.Password)
	return nil, 1
}

func checkUsernameCreate(user User) (int) {
	db := getConnection()
	row := db.QueryRow("SELECT username FROM users WHERE username = $1", user.Username)

	var username_finded string
	err := row.Scan(&username_finded)
	if (username_finded == "") {
		return 0
	}
	
	if err != nil {
		return 1
	}
	return 1
}

func checkEmailCreate(user User) (int) {
	db := getConnection()
	row := db.QueryRow("SELECT username FROM users WHERE email = $1", user.Email)

	var email_finded string
	err := row.Scan(&email_finded)
	if (email_finded == "") {
		return 0
	}
	
	if err != nil {
		return 1
	}
	return 1
}

func loginProcess(email string, password string) (int) {
	db := getConnection()
	row := db.QueryRow("SELECT username, email, firstname, lastname, password, country FROM users WHERE email = $1", email)

	var user_finded User
	err := row.Scan(&user_finded.Username, &user_finded.Email, &user_finded.Firstname, &user_finded.Lastname, &user_finded.Password, &user_finded.Country)
	if(user_finded.Password == "") {
		return 2
	}
	
	user_finded_password := []byte(user_finded.Password)
	password_byte_login := []byte(password)
	err = bcrypt.CompareHashAndPassword(user_finded_password, password_byte_login)
	if(err == nil) {
		user_logged = user_finded
		return 1
	} else {
		return 3;
	}

	if err != nil {
		return 0
	}
	return 0
}

func checkUserLogged() (bool) {
	if (user_logged.Username != "") {
		return true
	} else {
		return false
	}
}

func logoutUser() (int) {
	if(checkUserLogged()) {
		user_logged.Username = ""
		user_logged.Firstname = ""
		user_logged.Lastname = ""
		user_logged.Password = ""
		user_logged.Email = ""
		user_logged.Country = ""
		return 1
	} else {
		return 0
	}
}