package main

import (
	"bufio"
	"fmt"
	"galleria.com/hash"
	"galleria.com/rand"
	"os"
	"strings"

	"galleria.com/models"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "galleria"
)

// User model
type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Orders []Order
}

// Order Model
type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
	hmac := hash.NewHMAC("my-secret-key")
	fmt.Println(hmac.Hash("this is my string to hash"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	us.DestructiveReset()

	user := models.User{
		Name:  "Kimi Ri",
		Email: "kimi@ri.com",
		Password: "password",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}

	// verify user has a remember and rememberhash
	fmt.Printf("%+v\n", user)
	if user.Remember == "" {
		panic("Invalid remember token")
	}

	// Now verify that we can lookup a user with that remember token
	user2, err := us.ByRemember(user.Remember)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *user2)

	//Update user
	//user.Name = "Updated User"
	//if err := us.Update(&user); err != nil {
	//	panic(err)
	//}

	//foundUser, err := us.ByEmail("kimi@ri.com")
	//if err != nil {
	//	panic(err)
	//}

	// Delete a user
	//if err := us.Delete(foundUser.ID); err != nil {
	//	panic(err)
	//}

	// Verify the user is deleted
	//_, err = us.ByID(foundUser.ID)
	//if err != models.ErrNotFound {
	//	panic("user was not deleted")
	//}
	// name, email := getInfo()

	// u := &User{
	// 	Name: name,
	// 	Email: email,
	// }
	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", u)
	// var user User

	// db.Preload("Orders").First(&user)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }

	// fmt.Println("Email:", user.Email)
	// fmt.Println("Number of orders: ", len(user.Orders))
	// fmt.Println("Orders:", user.Orders)
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	})
	if db.Error != nil {
		panic(db.Error)
	}
}
