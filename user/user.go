package user

import (
	"github.com/sam-peets/micro/db"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string   `bson:"username"`
	Uid      uint32   `bson:"uid"`
	Hash     [32]byte `bson:"hash"`
}

// func createUser(payload User) (*User, error) {
// todo
// }

func CreateUser(user *User) error {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	users, err := connection.GetCollection("users")
	if err != nil {
		return err
	}

	_, err = users.InsertOne(context, user)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(uid uint32) (*User, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	users, err := connection.GetCollection("users")
	if err != nil {
		return nil, err
	}

	res := users.FindOne(context, bson.M{"uid": uid})
	var user User
	err = res.Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
