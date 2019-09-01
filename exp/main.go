package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"sinistra/lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "lenslocked"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	// Create a user
	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com"}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	// NOTE: You may need to update the query code a bit as well
	foundUser, err := us.ByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

	// Update the call to ByID to instead be ByEmail
	foundUser, err = us.ByEmail("michael@dundermifflin.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)

}
