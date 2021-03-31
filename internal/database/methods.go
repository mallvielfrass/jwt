package database

import (
	"errors"
	"fmt"
)

//AddUser
func (db Database) CreateUser(login, hash string) error {
	u := User{
		Login: login,
		Hash:  hash,
	}
	if result := db.db.Create(&u); result.Error != nil {
		// 	fmt.Printf("conf with login %s is exists", u.Login)
		return errors.New("login is exists")
	}
	return nil
}
func (db Database) GetUser(login string) (User, error) {
	fmt.Println("add user")
	var st User
	if err := db.db.Where("login = ?", login).First(&st).Error; err != nil {
		fmt.Println(err) // error handling...
		return st, errors.New("session is not exist")
	}
	return st, nil
}
func (db Database) GetAllUsers() []User {
	var users []User
	db.db.Where("").Find(&users)
	return users
}
func (db Database) CreateUserTable() {

	err := db.db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
	}
}

func (db Database) CreateSessionTable() {

	err := db.db.AutoMigrate(&SessionTable{})
	if err != nil {
		fmt.Println(err)
	}
}

func (db Database) CreateSession(login, session string, expiry int) error {
	s := SessionTable{
		Login:   login,
		Session: session,
		Expiry:  expiry,
	}
	if result := db.db.Create(&s); result.Error != nil {
		return errors.New("session is exists")
	}
	return nil
}
func (db Database) SearchSession(login, session string) (SessionTable, error) {
	var st SessionTable

	if err := db.db.Where("login = ? and session = ? ", login, session).First(&st).Error; err != nil {
		fmt.Println(err) // error handling...
		return st, errors.New("session is not exist")
	}
	//fmt.Println(st)
	return st, nil
}
