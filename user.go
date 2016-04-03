package main

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"log"
	"time"
)

type User struct {
	Email              string
	HashedPassword     string `datastore:",noindex"`
	SubscriptionExpiry time.Time
	Trial              bool

	CreatedAt time.Time
}

const PASSWORD_COST int = 12
const TWO_WEEKS time.Duration = time.Hour * 24 * 14

func (u *User) Save() error {
	client := DatastoreClient
	u.CreatedAt = time.Now()
	ctx := context.TODO()
	key := datastore.NewKey(ctx, "User", u.Email, 0, nil)
	key, err := client.Put(ctx, key, u)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) AddTimePeriod(duration time.Duration) {
	u.SubscriptionExpiry = time.Now().Add(duration)
}

func (u *User) StartTrial() {
	u.AddTimePeriod(TWO_WEEKS)
	u.Trial = true
}

func (u *User) Get() error {
	client := DatastoreClient
	key := datastore.NewKey(Ctx, "User", u.Email, 0, nil)
	err := client.Get(Ctx, key, u)

	if err != nil {
		return err
	}
	return nil
}

func (u *User) SetPassword(password string) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), PASSWORD_COST)

	if err != nil {
		log.Fatal(err)
	}

	u.HashedPassword = string(pw)
}

func (u *User) CorrectPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)) == nil
}

// func UserFromEmail(email string) *User {
// 	ctx := context.Background()
// 	client := DatastoreClient
//
// 	query := datastore.NewQuery("User").
// 		Filter("email =", email)
// 	it := client.Run(ctx, query)
//
// 	var user User
// 	for {
// 		_, err := it.Next(&user)
// 		if err == datastore.Done {
// 			break
// 		}
// 		if err != nil {
// 			log.Fatalf("Error fetching next task: %v", err)
// 		}
// 	}
// 	return &user
// }

// func UserFromId(id uint64) *User {
// 	u := NewUser()
// 	return u.Load(id)
// }

func NewUserFromEmail(email string) *User {
	u := NewUser()
	u.Email = email
	return u
}

func NewUserFromTemp(tu *TempUser) *User {
	u := NewUser()
	u.Email = tu.Email
	return u
}

func NewUser() *User {
	u := new(User)
	return u
}
