package auth

import (
	"crypto/sha256"
	"errors"
	"math/rand"
	"time"

	"github.com/sam-peets/micro/db"
	"github.com/sam-peets/micro/user"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	Sid     uint64
	Uid     uint64
	Expires time.Time
}

type UserPayload struct {
	Username string
	Password string
}

func NewUser(payload *UserPayload) (*user.User, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	users, err := connection.GetCollection("users")
	if err != nil {
		return nil, err
	}

	res := users.FindOne(context, bson.M{"username": payload.Username})
	if !errors.Is(res.Err(), mongo.ErrNoDocuments) {
		return nil, errors.New("user already exists")
	}

	hashed_password := sha256.Sum256([]byte(payload.Password))
	uid := rand.Uint32()
	for _, err := user.GetUser(uid); err == nil; uid = rand.Uint32() {
	}

	u := user.User{
		Hash:     hashed_password,
		Uid:      uid,
		Username: payload.Username,
	}

	_, err = users.InsertOne(context, u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func Auth(payload *UserPayload) (*Session, error) {
	connection := db.GetClient()
	context, cancel := connection.Context()
	defer cancel()

	users, err := connection.GetCollection("users")
	if err != nil {
		return nil, err
	}

	hashed_password := sha256.Sum256([]byte(payload.Password))
	var u user.User
	err = users.FindOne(context, bson.M{"username": payload.Username, "hash": hashed_password}).Decode(&u)
	if err != nil {
		return nil, err
	}

}
